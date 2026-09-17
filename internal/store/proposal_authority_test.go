package store

import (
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
)

func TestStoryProposalStorePersistsOnlyPendingAuthority(t *testing.T) {
	st := NewStore(t.TempDir())
	proposal := domain.StoryProposalV1{
		ID: "PROPOSAL-P001", Statement: "凤凰具有火属性。", Authority: domain.StoryProposalAuthority,
		Status: domain.StoryProposalStatusPending, SourceRefs: []string{"controlled-test"},
	}
	first, err := st.AppendStoryProposal(proposal)
	if err != nil {
		t.Fatal(err)
	}
	second, err := st.AppendStoryProposal(proposal)
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := st.LoadStoryProposalRegistry()
	if err != nil {
		t.Fatal(err)
	}
	if len(first.Proposals) != 1 || first.Digest != second.Digest || loaded == nil || loaded.Digest != first.Digest || loaded.Proposals[0].AcceptedCanon || loaded.Proposals[0].HumanApproved {
		t.Fatalf("proposal persistence changed authority: first=%+v second=%+v loaded=%+v", first, second, loaded)
	}
}
