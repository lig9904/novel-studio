package tools

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/rag"
	"github.com/chenhongyang/novel-studio/internal/store"
)

func proposalAuthorityTestStore(t *testing.T) *store.Store {
	t.Helper()
	st := store.NewStore(t.TempDir())
	_, err := st.AppendStoryProposal(domain.StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: domain.StoryProposalAuthority,
		Status: domain.StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	})
	if err != nil {
		t.Fatal(err)
	}
	return st
}

func TestProposalContextIsDiscussableButNotFactAuthority(t *testing.T) {
	st := proposalAuthorityTestStore(t)
	ctx, err := StoryProposalContext(st)
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := json.Marshal(ctx)
	for _, required := range []string{"PROPOSAL-P001", "凤凰具有火属性", "PROPOSAL", "PENDING", "accepted_canon\":false", "human_approved\":false"} {
		if !strings.Contains(string(raw), required) {
			t.Fatalf("proposal context missing %q: %s", required, raw)
		}
	}
	if err := ValidateNoPendingProposalClaims(st, "positive", map[string]any{"goal": "凤凰具有火属性。"}); err == nil {
		t.Fatal("pending proposal became a positive fact")
	}
	for _, payload := range []any{"凤凰具备火属性", map[string]any{"character": "凤凰", "attribute": "火"}} {
		if err := ValidateNoPendingProposalClaims(st, "paraphrase", payload); err == nil {
			t.Fatalf("proposal paraphrase bypassed Host gate: %#v", payload)
		}
	}
	if err := ValidateNoPendingProposalClaims(st, "negative", map[string]any{"forbidden_moves": []string{"不得断言凤凰具有火属性。"}}); err != nil {
		t.Fatalf("negative boundary was rejected: %v", err)
	}
	if err := ValidateNoPendingProposalClaims(st, "source metadata", map[string]any{"source_bindings": []any{map[string]any{"authority": "PROPOSAL", "usable_facts": []string{"凤凰具有火属性。"}}}}); err != nil {
		t.Fatalf("proposal provenance metadata was mistaken for a story fact: %v", err)
	}
}

func TestProposalRAGPersistsButCannotEnterFactRecall(t *testing.T) {
	st := proposalAuthorityTestStore(t)
	proposal := rag.RehashChunk(domain.RAGChunk{
		ID: "proposal-p001", SourcePath: "meta/story_proposals.json", SourceKind: rag.ProposalSourceKind,
		Text: "凤凰具有火属性。", Summary: "候选凤凰能力",
		Metadata: map[string]any{"proposal_id": "PROPOSAL-P001", "authority": "PROPOSAL", "accepted_canon": false, "human_approved": false},
	})
	if err := UpsertRAGChunks(context.Background(), st, nil, nil, []domain.RAGChunk{proposal}, domain.RAGIndexConfig{}); err != nil {
		t.Fatal(err)
	}
	index, err := st.RAG.LoadIndexStateReadOnly()
	if err != nil || index == nil || len(index.Chunks) != 1 {
		t.Fatalf("proposal was not durably retained: index=%+v err=%v", index, err)
	}
	if facts := activeRAGFactChunks(index); len(facts) != 0 {
		t.Fatalf("proposal entered ordinary fact corpus: %+v", facts)
	}

	contaminated := rag.RehashChunk(domain.RAGChunk{ID: "fact-p001", SourcePath: "summaries/01.json", SourceKind: "chapter_summary_facts", Text: "凤凰具有火属性。"})
	if err := UpsertRAGChunks(context.Background(), st, nil, nil, []domain.RAGChunk{contaminated}, domain.RAGIndexConfig{}); err == nil {
		t.Fatal("pending proposal was relabelled as a RAG fact")
	}
}

func TestProposalRAGRecoveryAndLegacyIndexCannotBypassAuthority(t *testing.T) {
	contaminated := rag.RehashChunk(domain.RAGChunk{ID: "fact-p001", SourcePath: "summaries/01.json", SourceKind: "chapter_summary_facts", Text: "凤凰的属性是火"})
	safe := rag.RehashChunk(domain.RAGChunk{ID: "safe", SourcePath: "summaries/02.json", SourceKind: "chapter_summary_facts", Text: "海风转向"})
	t.Run("pending queue", func(t *testing.T) {
		st := proposalAuthorityTestStore(t)
		if err := st.RAG.SavePendingUpserts(domain.RAGPendingUpserts{Chunks: []domain.RAGChunk{contaminated}}); err != nil {
			t.Fatal(err)
		}
		if err := UpsertRAGChunks(context.Background(), st, nil, nil, []domain.RAGChunk{safe}, domain.RAGIndexConfig{}); err == nil {
			t.Fatal("pending contaminated fact reentered RAG")
		}
	})
	t.Run("legacy active index", func(t *testing.T) {
		st := proposalAuthorityTestStore(t)
		if err := st.RAG.SaveIndexState(domain.RAGIndexState{Chunks: []domain.RAGChunk{contaminated}, ChunkHashes: []string{contaminated.Hash}}); err != nil {
			t.Fatal(err)
		}
		if err := UpsertRAGChunks(context.Background(), st, nil, nil, []domain.RAGChunk{safe}, domain.RAGIndexConfig{}); err == nil {
			t.Fatal("legacy contaminated fact survived current authority validation")
		}
	})
}

func TestSubmitStoryProposalCannotSelfApprove(t *testing.T) {
	st := store.NewStore(t.TempDir())
	tool := NewSubmitStoryProposalTool(st)
	raw, err := tool.Execute(context.Background(), json.RawMessage(`{"id":"PROPOSAL-P001","statement":"凤凰具有火属性。","source_refs":["controlled-test"]}`))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"accepted_canon":false`) || !strings.Contains(string(raw), `"human_approved":false`) {
		t.Fatalf("tool response did not preserve authority: %s", raw)
	}
	if _, err := tool.Execute(context.Background(), json.RawMessage(`{"id":"PROPOSAL-P002","statement":"九尾狐具有冰属性。","source_refs":["x"],"human_approved":true}`)); err == nil {
		t.Fatal("model supplied human approval through proposal tool")
	}
	if _, err := tool.Execute(context.Background(), json.RawMessage(`{"id":"PROPOSAL-P002","statement":"凤凰具有火属性。","source_refs":["x"],"subject_terms":["无关"]}`)); err == nil {
		t.Fatal("model supplied proposal match identity through proposal tool")
	}
}

func TestPendingProposalIsRejectedAtFactBearingToolBoundaries(t *testing.T) {
	st := proposalAuthorityTestStore(t)
	claim := "凤凰具有火属性。"
	cases := []struct {
		name string
		run  func() error
	}{
		{"foundation", func() error {
			_, err := NewSaveFoundationTool(st).Execute(context.Background(), json.RawMessage(`{"type":"premise","content":"凤凰具有火属性。","scale":"short"}`))
			return err
		}},
		{"volume_summary", func() error {
			_, err := NewSaveVolumeSummaryTool(st).Execute(context.Background(), json.RawMessage(`{"volume":1,"title":"候选","summary":"凤凰具有火属性。","key_events":["候选"]}`))
			return err
		}},
		{"arc_summary", func() error {
			_, err := NewSaveArcSummaryTool(st).Execute(context.Background(), json.RawMessage(`{"volume":1,"arc":1,"title":"候选","summary":"凤凰具有火属性。","key_events":["候选"],"character_snapshots":[]}`))
			return err
		}},
		{"world_simulation", func() error {
			_, err := NewSimulateChapterWorldTool(st).Execute(context.Background(), json.RawMessage(`{"chapter":1,"time_window":"凤凰具有火属性。","character_decisions":[],"sources":[],"finalize":false}`))
			return err
		}},
		{"world_tick", func() error {
			_, err := NewSaveWorldTickTool(st).Execute(context.Background(), json.RawMessage(`{"volume":1,"arc":1,"through_chapter":0,"events":[{"summary":"凤凰具备火属性"}]}`))
			return err
		}},
		{"chapter_commit", func() error {
			_, err := NewCommitChapterTool(st).Execute(context.Background(), json.RawMessage(`{"chapter":1,"summary":"凤凰的属性是火","characters":[],"key_events":[]}`))
			return err
		}},
		{"character_decision", func() error {
			observation := domain.CharacterObservationPacket{Version: domain.CharacterObservationVersion}
			_, err := NewSubmitCharacterDecisionTool(st, observation).Execute(context.Background(), json.RawMessage(`{"location":"原地","current_goal":"凤凰具有火属性。","pressure":"无","available_options":["等待","观察"],"decision":"等待","decision_reason":"保持边界","intended_action":"等待","action_duration":"1分钟","knowledge_refs":[]}`))
			return err
		}},
		{"world_arbitration", func() error {
			proposal := domain.CharacterDecisionProposal{AgentID: "agent", Character: "角色", Round: 1}
			tool := NewResolveChapterWorldTool(st, domain.WorldStimulusPacket{}, domain.CharacterAgentActivation{}, []domain.CharacterDecisionProposal{proposal}, "protocol", nil, 1)
			_, err := tool.Execute(context.Background(), json.RawMessage(`{"time_window":"凤凰具有火属性。","resolutions":[],"hard_contract_status":"feasible","hard_contract_conflicts":[],"protagonist_projection":{"protagonist":"角色","observable_effects":[],"hidden_pressures":[],"available_options":[],"chosen_decision":"等待","decision_reason":"边界","plan_constraints":[],"causal_chain":[]},"finalized":false}`))
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.run()
			if err == nil || !strings.Contains(err.Error(), "PROPOSAL-P001") {
				t.Fatalf("%s did not reject %q: %v", tc.name, claim, err)
			}
		})
	}
}

func TestArchitectOnlyProposalCapabilityDoesNotLeakThroughSharedContext(t *testing.T) {
	st := proposalAuthorityTestStore(t)
	shared, err := NewContextTool(st, References{}, "").Execute(context.Background(), json.RawMessage(`{"chapter":0}`))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(shared), `PROPOSAL-P001`) || strings.Contains(string(shared), `proposal_context`) {
		t.Fatalf("shared novel_context leaked proposal material: %s", shared)
	}
	raw, err := NewListStoryProposalsTool(st).Execute(context.Background(), json.RawMessage(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"PROPOSAL-P001"`) || !strings.Contains(string(raw), `"accepted_canon":false`) {
		t.Fatalf("dedicated proposal capability lost authority labels: %s", raw)
	}
}
