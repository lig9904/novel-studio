package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/store"
)

func TestP06ExistingProjectAllWorkspaceRejectsNewLiveProposalRegistry(t *testing.T) {
	live := t.TempDir()
	workspace := t.TempDir()
	foundation, err := pipelineProjectAllFoundationSnapshotRoot(live)
	if err != nil {
		t.Fatal(err)
	}
	ragRoot, err := pipelineProjectAllRAGSnapshotRoot(workspace)
	if err != nil {
		t.Fatal(err)
	}
	manifest := pipelineProjectAllWorkspaceManifest{
		Version: "project-all-workspace.v3", GenerationID: "pg2_p06", SourceOutput: filepath.Clean(live),
		BaseChapter: 0, Workspace: filepath.Clean(workspace), IsolatedWrites: true,
		FoundationSnapshotRoot: foundation, RAGSnapshotRoot: ragRoot,
	}
	raw, _ := json.Marshal(manifest)
	path := filepath.Join(workspace, filepath.FromSlash(pipelineProjectAllWorkspaceManifestPath))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := validatePipelineProjectAllWorkspaceManifest(workspace, live, "pg2_p06", 0); err != nil {
		t.Fatalf("baseline workspace invalid: %v", err)
	}
	if _, err := store.NewStore(live).AppendStoryProposal(domain.StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: domain.StoryProposalAuthority,
		Status: domain.StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	}); err != nil {
		t.Fatal(err)
	}
	if err := validatePipelineProjectAllWorkspaceManifest(workspace, live, "pg2_p06", 0); err == nil || !strings.Contains(err.Error(), "does not match") {
		t.Fatalf("stale workspace reused after proposal registry appeared: %v", err)
	}
}

func TestP06WorkspaceManifestAuthenticatesShadowProposalRegistry(t *testing.T) {
	proposal := domain.StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: domain.StoryProposalAuthority,
		Status: domain.StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	}
	for _, mutation := range []string{"missing", "replaced"} {
		t.Run(mutation, func(t *testing.T) {
			live, workspace := t.TempDir(), t.TempDir()
			if _, err := store.NewStore(live).AppendStoryProposal(proposal); err != nil {
				t.Fatal(err)
			}
			if _, err := store.NewStore(workspace).AppendStoryProposal(proposal); err != nil {
				t.Fatal(err)
			}
			foundation, err := pipelineProjectAllFoundationSnapshotRoot(live)
			if err != nil {
				t.Fatal(err)
			}
			ragRoot, err := pipelineProjectAllRAGSnapshotRoot(workspace)
			if err != nil {
				t.Fatal(err)
			}
			proposalDigest, err := pipelineProjectAllProposalRegistryDigest(workspace)
			if err != nil {
				t.Fatal(err)
			}
			manifest := pipelineProjectAllWorkspaceManifest{
				Version: "project-all-workspace.v3", GenerationID: "pg2_p06", SourceOutput: filepath.Clean(live),
				BaseChapter: 0, Workspace: filepath.Clean(workspace), IsolatedWrites: true,
				FoundationSnapshotRoot: foundation, RAGSnapshotRoot: ragRoot, ProposalRegistryDigest: proposalDigest,
			}
			raw, _ := json.Marshal(manifest)
			path := filepath.Join(workspace, filepath.FromSlash(pipelineProjectAllWorkspaceManifestPath))
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, raw, 0o644); err != nil {
				t.Fatal(err)
			}
			if err := validatePipelineProjectAllWorkspaceManifest(workspace, live, "pg2_p06", 0); err != nil {
				t.Fatalf("baseline workspace invalid: %v", err)
			}
			if err := os.Remove(filepath.Join(workspace, filepath.FromSlash(store.StoryProposalRegistryPath))); err != nil {
				t.Fatal(err)
			}
			if mutation == "replaced" {
				if _, err := store.NewStore(workspace).AppendStoryProposal(domain.StoryProposalV1{
					ID: "PROPOSAL-P002", Statement: "九尾狐具有冰属性。", Authority: domain.StoryProposalAuthority,
					Status: domain.StoryProposalStatusPending, SourceRefs: []string{"replacement"},
				}); err != nil {
					t.Fatal(err)
				}
			}
			if err := validatePipelineProjectAllWorkspaceManifest(workspace, live, "pg2_p06", 0); err == nil {
				t.Fatalf("%s shadow proposal registry passed frozen manifest", mutation)
			}
		})
	}
}
