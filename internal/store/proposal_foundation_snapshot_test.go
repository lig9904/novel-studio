package store

import (
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
)

func TestProjectAllFoundationSnapshotBindsProposalRegistryWhenPresent(t *testing.T) {
	dir := t.TempDir()
	_, before, err := CaptureProjectAllFoundationSnapshot(dir)
	if err != nil {
		t.Fatal(err)
	}
	st := NewStore(dir)
	if _, err := st.AppendStoryProposal(domain.StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: domain.StoryProposalAuthority,
		Status: domain.StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	}); err != nil {
		t.Fatal(err)
	}
	snapshot, after, err := CaptureProjectAllFoundationSnapshot(dir)
	if err != nil {
		t.Fatal(err)
	}
	if before == after || snapshot.Artifacts[StoryProposalRegistryPath] == "" {
		t.Fatalf("proposal registry did not change foundation identity: before=%s after=%s snapshot=%+v", before, after, snapshot)
	}
}
