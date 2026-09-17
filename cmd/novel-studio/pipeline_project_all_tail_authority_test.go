package main

import (
	"strings"
	"testing"

	"github.com/chenhongyang/novel-studio/internal/domain"
)

func sealProjectAllTailTestGeneration(
	t *testing.T,
	generation domain.PlanningGenerationV2,
	bundles []domain.ProjectedChapterBundle,
	registry *domain.ObligationRegistryV2,
) domain.PlanningGenerationV2 {
	t.Helper()
	var err error
	registry.RegistryRoot, err = domain.ComputeObligationRegistryV2Root(*registry)
	if err != nil {
		t.Fatal(err)
	}
	generation.Status = domain.PlanningGenerationSealedV2
	generation.ProjectedChapterCount = len(bundles)
	generation.ChainHeadRoot = bundles[0].BundleDigest
	generation.ChainTailRoot = bundles[len(bundles)-1].BundleDigest
	generation.ObligationRegistryRoot = registry.RegistryRoot
	generation.SealedAt = "2026-09-17T00:00:00Z"
	generation.GenerationDigest, err = domain.ComputePlanningGenerationV2Digest(generation)
	if err != nil {
		t.Fatal(err)
	}
	return generation
}

func TestPipelineProjectAllTailAuthorityMatrix(t *testing.T) {
	t.Run("first bundle uses derived generation genesis", func(t *testing.T) {
		generation, registry := projectAllCmdTestGenerationAndRegistry(t, 1)
		previous, preState, err := pipelineProjectAllTail(generation, nil)
		if err != nil {
			t.Fatal(err)
		}
		genesis, err := domain.DeriveProjectedChainGenesisV2(generation)
		if err != nil {
			t.Fatal(err)
		}
		if previous != genesis || preState != generation.BaseStateRoot {
			t.Fatalf("first tail=(%s,%s) want genesis/base=(%s,%s)", previous, preState, genesis, generation.BaseStateRoot)
		}
		emptyPrevious, emptyPreState, err := pipelineProjectAllTail(generation, []domain.ProjectedChapterBundle{})
		if err != nil || emptyPrevious != previous || emptyPreState != preState {
			t.Fatalf("empty slice diverged from nil tail: previous=%s pre=%s err=%v", emptyPrevious, emptyPreState, err)
		}
		artifacts, outline := projectAllCmdTestArtifacts(t, generation.GenerationID, 1)
		projectAllCmdTestBindPlanningContext(t, artifacts, generation, nil, registry, 1)
		bundle, nextRegistry, err := buildPipelineProjectedChapterBundle(generation, outline, previous, preState, artifacts, registry)
		if err != nil {
			t.Fatal(err)
		}
		bundles := []domain.ProjectedChapterBundle{bundle}
		generation = sealProjectAllTailTestGeneration(t, generation, bundles, &nextRegistry)
		if err := domain.ValidateProjectedChapterBundleChain(generation, bundles, nextRegistry); err != nil {
			t.Fatalf("first bundle did not consume the production genesis: %v", err)
		}
	})

	t.Run("second bundle uses chapter one digest", func(t *testing.T) {
		generation, registry := projectAllCmdTestGenerationAndRegistry(t, 2)
		previous, preState, err := pipelineProjectAllTail(generation, nil)
		if err != nil {
			t.Fatal(err)
		}
		firstArtifacts, firstOutline := projectAllCmdTestArtifacts(t, generation.GenerationID, 1)
		projectAllCmdTestBindPlanningContext(t, firstArtifacts, generation, nil, registry, 1)
		first, registry, err := buildPipelineProjectedChapterBundle(generation, firstOutline, previous, preState, firstArtifacts, registry)
		if err != nil {
			t.Fatal(err)
		}
		previous, preState, err = pipelineProjectAllTail(generation, []domain.ProjectedChapterBundle{first})
		if err != nil {
			t.Fatal(err)
		}
		if previous != first.BundleDigest || preState != first.ProjectedPostStateRoot {
			t.Fatalf("second tail lost chapter one authority: previous=%s pre=%s", previous, preState)
		}
		secondArtifacts, secondOutline := projectAllCmdTestArtifacts(t, generation.GenerationID, 2)
		projectAllCmdTestBindPlanningContext(t, secondArtifacts, generation, []domain.ProjectedChapterBundle{first}, registry, 2)
		second, registry, err := buildPipelineProjectedChapterBundle(generation, secondOutline, previous, preState, secondArtifacts, registry)
		if err != nil {
			t.Fatal(err)
		}
		bundles := []domain.ProjectedChapterBundle{first, second}
		generation = sealProjectAllTailTestGeneration(t, generation, bundles, &registry)
		if err := domain.ValidateProjectedChapterBundleChain(generation, bundles, registry); err != nil {
			t.Fatalf("two-bundle production tail chain failed: %v", err)
		}
	})

	t.Run("empty digest is rejected", func(t *testing.T) {
		generation, registry := projectAllCmdTestGenerationAndRegistry(t, 1)
		artifacts, outline := projectAllCmdTestArtifacts(t, generation.GenerationID, 1)
		projectAllCmdTestBindPlanningContext(t, artifacts, generation, nil, registry, 1)
		_, _, err := buildPipelineProjectedChapterBundle(generation, outline, "", generation.BaseStateRoot, artifacts, registry)
		if err == nil || !strings.Contains(err.Error(), "previous_bundle_digest") {
			t.Fatalf("empty predecessor digest was not rejected: %v", err)
		}
	})

	for _, tc := range []struct {
		name     string
		previous func(domain.PlanningGenerationV2) string
	}{
		{name: "formatted fake digest", previous: func(domain.PlanningGenerationV2) string { return projectAllCmdTestDigest("not-the-real-tail") }},
		{name: "wrong generation genesis", previous: func(generation domain.PlanningGenerationV2) string {
			generation.GenerationID = "pg2_wrong_generation"
			genesis, err := domain.DeriveProjectedChainGenesisV2(generation)
			if err != nil {
				t.Fatal(err)
			}
			return genesis
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			generation, registry := projectAllCmdTestGenerationAndRegistry(t, 1)
			artifacts, outline := projectAllCmdTestArtifacts(t, generation.GenerationID, 1)
			projectAllCmdTestBindPlanningContext(t, artifacts, generation, nil, registry, 1)
			bundle, nextRegistry, err := buildPipelineProjectedChapterBundle(generation, outline, tc.previous(generation), generation.BaseStateRoot, artifacts, registry)
			if err != nil {
				t.Fatalf("well-formed wrong predecessor should reach chain authority check: %v", err)
			}
			bundles := []domain.ProjectedChapterBundle{bundle}
			generation = sealProjectAllTailTestGeneration(t, generation, bundles, &nextRegistry)
			if err := domain.ValidateProjectedChapterBundleChain(generation, bundles, nextRegistry); err == nil || !strings.Contains(err.Error(), "previous_bundle_digest does not match genesis") {
				t.Fatalf("wrong predecessor authority passed chain validation: %v", err)
			}
		})
	}
}
