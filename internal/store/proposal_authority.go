package store

import (
	"fmt"
	"os"

	"github.com/chenhongyang/novel-studio/internal/domain"
)

const StoryProposalRegistryPath = "meta/story_proposals.json"

func (s *Store) SaveStoryProposalRegistry(registry domain.StoryProposalRegistryV1) error {
	finalized, err := domain.FinalizeStoryProposalRegistryV1(registry)
	if err != nil {
		return err
	}
	return s.Progress.io.WriteJSON(StoryProposalRegistryPath, finalized)
}

func (s *Store) LoadStoryProposalRegistry() (*domain.StoryProposalRegistryV1, error) {
	var registry domain.StoryProposalRegistryV1
	if err := s.Progress.io.ReadJSON(StoryProposalRegistryPath, &registry); err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	if err := domain.ValidateStoryProposalRegistryV1(registry); err != nil {
		return nil, fmt.Errorf("invalid story proposal registry: %w", err)
	}
	return &registry, nil
}

func (s *Store) AppendStoryProposal(proposal domain.StoryProposalV1) (*domain.StoryProposalRegistryV1, error) {
	single, err := domain.FinalizeStoryProposalRegistryV1(domain.StoryProposalRegistryV1{
		Policy: domain.StoryProposalPolicyV1, Proposals: []domain.StoryProposalV1{proposal},
	})
	if err != nil {
		return nil, err
	}
	proposal = single.Proposals[0]
	var result domain.StoryProposalRegistryV1
	err = s.Progress.io.WithWriteLock(func() error {
		registry := domain.StoryProposalRegistryV1{Policy: domain.StoryProposalPolicyV1}
		if err := s.Progress.io.ReadJSONUnlocked(StoryProposalRegistryPath, &registry); err != nil && !os.IsNotExist(err) {
			return err
		}
		for _, existing := range registry.Proposals {
			if existing.ID == proposal.ID {
				if existing.Digest == proposal.Digest && existing.Statement == proposal.Statement {
					result = registry
					return nil
				}
				return fmt.Errorf("story proposal %q already exists with different content", proposal.ID)
			}
		}
		registry.Proposals = append(registry.Proposals, proposal)
		registry.Digest = ""
		finalized, err := domain.FinalizeStoryProposalRegistryV1(registry)
		if err != nil {
			return err
		}
		if err := s.Progress.io.WriteJSONUnlocked(StoryProposalRegistryPath, finalized); err != nil {
			return err
		}
		result = finalized
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
