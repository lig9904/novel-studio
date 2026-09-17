package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/errs"
	"github.com/chenhongyang/novel-studio/internal/rag"
	"github.com/chenhongyang/novel-studio/internal/store"
	"github.com/voocel/agentcore/schema"
)

// SubmitStoryProposalTool persists candidate material without any path to
// approval or Canon. The fixed authority/status fields are host-owned.
type SubmitStoryProposalTool struct{ store *store.Store }

func NewSubmitStoryProposalTool(st *store.Store) *SubmitStoryProposalTool {
	return &SubmitStoryProposalTool{store: st}
}

func (t *SubmitStoryProposalTool) Name() string                         { return "submit_story_proposal" }
func (t *SubmitStoryProposalTool) Label() string                        { return "提交候选设定" }
func (t *SubmitStoryProposalTool) ReadOnly(json.RawMessage) bool        { return false }
func (t *SubmitStoryProposalTool) ConcurrencySafe(json.RawMessage) bool { return false }
func (t *SubmitStoryProposalTool) Description() string {
	return "保存可讨论的候选设定。结果固定为 PROPOSAL/PENDING/NOT ACCEPTED CANON/NOT HUMAN APPROVED；本工具不能批准或写入正史。"
}
func (t *SubmitStoryProposalTool) Schema() map[string]any {
	return schema.Object(
		schema.Property("id", schema.String("稳定 Proposal ID，例如 PROPOSAL-P001")).Required(),
		schema.Property("statement", schema.String("单一、明确、可审阅的候选设定陈述")).Required(),
		schema.Property("source_refs", schema.Array("候选来源或讨论依据；不表示事实权威", schema.String("source ref"))).Required(),
	)
}
func (t *SubmitStoryProposalTool) Execute(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
	if t.store == nil {
		return nil, fmt.Errorf("story proposal store is unavailable")
	}
	var input struct {
		ID             string    `json:"id"`
		Statement      string    `json:"statement"`
		SourceRefs     []string  `json:"source_refs"`
		Authority      string    `json:"authority"`
		Status         string    `json:"status"`
		AcceptedCanon  *bool     `json:"accepted_canon"`
		HumanApproved  *bool     `json:"human_approved"`
		SubjectTerms   *[]string `json:"subject_terms"`
		PredicateTerms *[]string `json:"predicate_terms"`
		ObjectTerms    *[]string `json:"object_terms"`
	}
	if err := unmarshalToolArgs(args, &input); err != nil {
		return nil, fmt.Errorf("invalid args: %w: %w", errs.ErrToolArgs, err)
	}
	if input.Authority != "" || input.Status != "" || input.AcceptedCanon != nil || input.HumanApproved != nil || input.SubjectTerms != nil || input.PredicateTerms != nil || input.ObjectTerms != nil {
		return nil, fmt.Errorf("proposal authority/status/approval fields are host-owned: %w", errs.ErrToolPrecondition)
	}
	proposal := domain.StoryProposalV1{
		ID: strings.TrimSpace(input.ID), Statement: strings.TrimSpace(input.Statement), SourceRefs: input.SourceRefs,
		Authority: domain.StoryProposalAuthority, Status: domain.StoryProposalStatusPending,
		AcceptedCanon: false, HumanApproved: false,
	}
	registry, err := t.store.AppendStoryProposal(proposal)
	if err != nil {
		return nil, err
	}
	for _, stored := range registry.Proposals {
		if stored.ID == proposal.ID {
			return json.Marshal(map[string]any{
				"saved": true, "id": stored.ID, "digest": stored.Digest,
				"authority": stored.Authority, "status": stored.Status,
				"accepted_canon": false, "human_approved": false,
			})
		}
	}
	return nil, fmt.Errorf("saved story proposal is missing from registry")
}

type ListStoryProposalsTool struct{ store *store.Store }

func NewListStoryProposalsTool(st *store.Store) *ListStoryProposalsTool {
	return &ListStoryProposalsTool{store: st}
}
func (t *ListStoryProposalsTool) Name() string                         { return "list_story_proposals" }
func (t *ListStoryProposalsTool) Label() string                        { return "读取候选设定" }
func (t *ListStoryProposalsTool) ReadOnly(json.RawMessage) bool        { return true }
func (t *ListStoryProposalsTool) ConcurrencySafe(json.RawMessage) bool { return true }
func (t *ListStoryProposalsTool) Description() string {
	return "读取明确隔离的候选设定供Architect讨论/比较；返回内容不具有Story Truth或Canon Authority。"
}
func (t *ListStoryProposalsTool) Schema() map[string]any { return schema.Object() }
func (t *ListStoryProposalsTool) Execute(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
	var empty struct{}
	if err := unmarshalToolArgs(args, &empty); err != nil {
		return nil, fmt.Errorf("invalid args: %w: %w", errs.ErrToolArgs, err)
	}
	value, err := StoryProposalContext(t.store)
	if err != nil {
		return nil, err
	}
	if value == nil {
		value = map[string]any{"policy": domain.StoryProposalPolicyV1, "proposals": []any{}, "accepted_canon": false, "human_approved": false}
	}
	return json.Marshal(value)
}

func StoryProposalContext(st *store.Store) (map[string]any, error) {
	if st == nil {
		return nil, nil
	}
	registry, err := st.LoadStoryProposalRegistry()
	if err != nil || registry == nil {
		return nil, err
	}
	return map[string]any{
		"policy": registry.Policy, "registry_digest": registry.Digest,
		"authority":      domain.StoryProposalAuthority,
		"accepted_canon": false, "human_approved": false,
		"usage_policy": "候选只可讨论、比较或进入隔离模拟；Presence/receipt/digest均不授予Story Truth。没有独立Host/Human approval时不得写入foundation、summary、Character/World、Plan、Bundle或Canon。",
		"proposals":    registry.Proposals,
	}, nil
}

// ValidateNoPendingProposalClaims rejects registered candidate statements on
// fact-bearing surfaces. Negative/source metadata is removed by the shared
// domain projection before matching.
func ValidateNoPendingProposalClaims(st *store.Store, label string, payload any) error {
	if st == nil {
		return nil
	}
	registry, err := st.LoadStoryProposalRegistry()
	if err != nil || registry == nil || len(registry.Proposals) == 0 {
		return err
	}
	for _, proposal := range registry.Proposals {
		if domain.StoryProposalMatchesValueV1(proposal, payload) {
			return fmt.Errorf("%s asserts pending proposal %s without authorized approval: %w", label, proposal.ID, errs.ErrToolPrecondition)
		}
	}
	return nil
}

func validateProposalRAGChunks(st *store.Store, chunks []domain.RAGChunk) error {
	if st == nil {
		return fmt.Errorf("proposal-aware RAG indexing requires a project store")
	}
	registry, err := st.LoadStoryProposalRegistry()
	if err != nil {
		return err
	}
	byID := map[string]domain.StoryProposalV1{}
	if registry != nil {
		for _, proposal := range registry.Proposals {
			byID[proposal.ID] = proposal
		}
	}
	for _, chunk := range chunks {
		if strings.EqualFold(strings.TrimSpace(chunk.SourceKind), rag.ProposalSourceKind) {
			id, _ := chunk.Metadata["proposal_id"].(string)
			authority, _ := chunk.Metadata["authority"].(string)
			accepted, _ := chunk.Metadata["accepted_canon"].(bool)
			approved, _ := chunk.Metadata["human_approved"].(bool)
			proposal, ok := byID[strings.TrimSpace(id)]
			if !ok || strings.TrimSpace(authority) != domain.StoryProposalAuthority || accepted || approved ||
				!domain.StoryProposalMatchesTextV1(proposal, chunk.Text+" "+chunk.Summary) {
				return fmt.Errorf("proposal RAG chunk %q lacks its exact pending registry authority", chunk.ID)
			}
			continue
		}
		if err := ValidateNoPendingProposalClaims(st, "RAG fact chunk "+chunk.ID, chunk); err != nil {
			return err
		}
	}
	return nil
}
