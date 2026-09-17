package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

const (
	StoryProposalPolicyV1       = "story-proposal.v1"
	StoryProposalAuthority      = "PROPOSAL"
	StoryProposalStatusPending  = "PENDING"
	SourceAuthorityProposal     = "PROPOSAL"
	SourceAuthorityReference    = "REFERENCE"
	SourceAuthoritySimulation   = "SIMULATION_FACT"
	SourceAuthorityPresentation = "DERIVED_PRESENTATION"
	SourceAuthorityCanon        = "ACCEPTED_CANON"
)

// StoryProposalV1 is durable candidate material, not a story fact. This
// protocol deliberately has no model-writable approved state: promotion needs
// a separate host/human approval contract that does not exist in v1.
type StoryProposalV1 struct {
	ID             string   `json:"id"`
	Statement      string   `json:"statement"`
	Authority      string   `json:"authority"`
	Status         string   `json:"status"`
	AcceptedCanon  bool     `json:"accepted_canon"`
	HumanApproved  bool     `json:"human_approved"`
	SourceRefs     []string `json:"source_refs"`
	SubjectTerms   []string `json:"subject_terms"`
	PredicateTerms []string `json:"predicate_terms"`
	ObjectTerms    []string `json:"object_terms"`
	Digest         string   `json:"digest"`
}

type StoryProposalRegistryV1 struct {
	Policy    string            `json:"policy"`
	Proposals []StoryProposalV1 `json:"proposals"`
	Digest    string            `json:"digest"`
}

func FinalizeStoryProposalRegistryV1(registry StoryProposalRegistryV1) (StoryProposalRegistryV1, error) {
	if registry.Policy == "" {
		registry.Policy = StoryProposalPolicyV1
	}
	if registry.Policy != StoryProposalPolicyV1 {
		return registry, fmt.Errorf("unsupported story proposal policy %q", registry.Policy)
	}
	registry.Proposals = append([]StoryProposalV1(nil), registry.Proposals...)
	seen := make(map[string]struct{}, len(registry.Proposals))
	for i := range registry.Proposals {
		proposal := &registry.Proposals[i]
		proposal.ID = strings.TrimSpace(proposal.ID)
		proposal.Statement = strings.TrimSpace(proposal.Statement)
		proposal.Authority = strings.TrimSpace(proposal.Authority)
		proposal.Status = strings.TrimSpace(proposal.Status)
		proposal.SourceRefs = compactStoryProposalStrings(proposal.SourceRefs)
		claimedSubjectTerms := compactStoryProposalStrings(proposal.SubjectTerms)
		claimedPredicateTerms := compactStoryProposalStrings(proposal.PredicateTerms)
		claimedObjectTerms := compactStoryProposalStrings(proposal.ObjectTerms)
		derivedSubjectTerms, derivedPredicateTerms, derivedObjectTerms, err := deriveStoryProposalClaimV1(proposal.Statement)
		if err != nil {
			return registry, fmt.Errorf("story proposal %q: %w", proposal.ID, err)
		}
		if (len(claimedSubjectTerms) > 0 && !reflect.DeepEqual(claimedSubjectTerms, derivedSubjectTerms)) ||
			(len(claimedPredicateTerms) > 0 && !reflect.DeepEqual(claimedPredicateTerms, derivedPredicateTerms)) ||
			(len(claimedObjectTerms) > 0 && !reflect.DeepEqual(claimedObjectTerms, derivedObjectTerms)) {
			return registry, fmt.Errorf("story proposal %q claim terms are host-derived and do not match statement", proposal.ID)
		}
		proposal.SubjectTerms = derivedSubjectTerms
		proposal.PredicateTerms = derivedPredicateTerms
		proposal.ObjectTerms = derivedObjectTerms
		if !strings.HasPrefix(proposal.ID, "PROPOSAL-") || len(proposal.ID) > 128 || proposal.Statement == "" || len([]rune(proposal.Statement)) > 2000 {
			return registry, fmt.Errorf("story proposal %d requires bounded id and statement", i)
		}
		if _, duplicate := seen[proposal.ID]; duplicate {
			return registry, fmt.Errorf("duplicate story proposal id %q", proposal.ID)
		}
		seen[proposal.ID] = struct{}{}
		if proposal.Authority != StoryProposalAuthority || proposal.Status != StoryProposalStatusPending || proposal.AcceptedCanon || proposal.HumanApproved {
			return registry, fmt.Errorf("story proposal %q must remain PROPOSAL/PENDING and unapproved", proposal.ID)
		}
		if len(proposal.SourceRefs) == 0 {
			return registry, fmt.Errorf("story proposal %q requires source_refs", proposal.ID)
		}
		if len(proposal.SubjectTerms) == 0 || len(proposal.ObjectTerms) == 0 {
			return registry, fmt.Errorf("story proposal %q requires subject_terms and object_terms", proposal.ID)
		}
		claimed := proposal.Digest
		proposal.Digest = ""
		proposal.Digest = storyProposalDigest(*proposal)
		if claimed != "" && claimed != proposal.Digest {
			return registry, fmt.Errorf("story proposal %q digest mismatch", proposal.ID)
		}
	}
	sort.Slice(registry.Proposals, func(i, j int) bool { return registry.Proposals[i].ID < registry.Proposals[j].ID })
	claimed := registry.Digest
	registry.Digest = ""
	registry.Digest = storyProposalDigest(registry)
	if claimed != "" && claimed != registry.Digest {
		return registry, fmt.Errorf("story proposal registry digest mismatch")
	}
	return registry, nil
}

func ValidateStoryProposalRegistryV1(registry StoryProposalRegistryV1) error {
	finalized, err := FinalizeStoryProposalRegistryV1(registry)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(finalized, registry) {
		return fmt.Errorf("story proposal registry is not finalized")
	}
	return nil
}

func ValidSourceAuthorityV2(authority string) bool {
	switch strings.TrimSpace(authority) {
	case SourceAuthorityProposal, SourceAuthorityReference, SourceAuthoritySimulation, SourceAuthorityPresentation, SourceAuthorityCanon:
		return true
	default:
		return false
	}
}

func validateProjectedProposalBindingsV2(bundle ProjectedChapterBundle) error {
	var claims []string
	for _, binding := range bundle.SourceBindings {
		if strings.TrimSpace(binding.Authority) != SourceAuthorityProposal {
			continue
		}
		claims = append(claims, binding.UsableFacts...)
	}
	if len(claims) == 0 {
		return nil
	}
	surface := map[string]any{
		"chapter_world_simulation": bundle.ChapterWorldSimulation,
		"chapter_plan":             proposalPositiveValueV1(bundle.ChapterPlan),
		"formal_world_simulation":  bundle.FormalWorldSimulation,
		"pov_plan":                 bundle.POVPlan,
		"hard_render_contract": map[string]any{
			"must_occur": bundle.HardRenderContract.MustOccur, "must_preserve": bundle.HardRenderContract.MustPreserve,
			"foreshadow_changes": bundle.HardRenderContract.ForeshadowChanges, "resource_changes": bundle.HardRenderContract.ResourceChanges,
			"relationship_changes": bundle.HardRenderContract.RelationshipChanges, "knowledge_changes": bundle.HardRenderContract.KnowledgeChanges,
		},
		"projected_delta": bundle.ProjectedDelta,
		"render_context":  proposalPositiveValueV1(bundle.RenderContext),
	}
	for _, claim := range claims {
		proposal := StoryProposalV1{Statement: claim}
		if StoryProposalMatchesValueV1(proposal, surface) {
			return fmt.Errorf("pending proposal claim appears on a fact-bearing bundle surface")
		}
	}
	return nil
}

// ProposalPositiveValueV1 removes source, negative and boundary metadata before
// proposal-claim checks. Mentioning a proposal in a do-not-use rule is not the
// same as asserting it as a story fact.
func ProposalPositiveValueV1(value any) any { return proposalPositiveValueV1(value) }

func proposalPositiveValueV1(value any) any {
	raw, err := json.Marshal(value)
	if err != nil {
		return value
	}
	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return value
	}
	return stripProposalNegativeFieldsV1(decoded)
}

func stripProposalNegativeFieldsV1(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		for key, item := range typed {
			lower := strings.ToLower(strings.TrimSpace(key))
			if lower == "proposal_context" || lower == "source_bindings" || lower == "source_refs" || lower == "exact_references" ||
				lower == "query_or_need" || lower == "do_not_use" || lower == "knowledge_boundary" || lower == "reveal_budget" ||
				lower == "must_not_occur" || lower == "pov_does_not_know" || lower == "rejected_options" ||
				strings.Contains(lower, "constraint") || strings.Contains(lower, "forbidden") {
				continue
			}
			out[key] = stripProposalNegativeFieldsV1(item)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for i := range typed {
			out[i] = stripProposalNegativeFieldsV1(typed[i])
		}
		return out
	default:
		return value
	}
}

func NormalizeProposalClaimTextV1(value string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || strings.ContainsRune("，。、“”‘’！？；：,.!?;:\"'()（）[]【】", r) {
			return -1
		}
		return unicode.ToLower(r)
	}, strings.TrimSpace(value))
}

func StoryProposalMatchesTextV1(proposal StoryProposalV1, value string) bool {
	proposal = storyProposalWithDerivedTermsV1(proposal)
	for _, boundary := range []string{"但是", "不过", "并且", "而且", "同时", "但", "且"} {
		value = strings.ReplaceAll(value, boundary, "。")
	}
	for _, unit := range strings.FieldsFunc(value, func(r rune) bool {
		return r == '\n' || strings.ContainsRune("。！？!?；;，,", r)
	}) {
		if storyProposalMatchesClaimUnitV1(proposal, unit) {
			return true
		}
	}
	return false
}

func StoryProposalMatchesValueV1(proposal StoryProposalV1, value any) bool {
	proposal = storyProposalWithDerivedTermsV1(proposal)
	value = proposalPositiveValueV1(value)
	var visit func(any) bool
	visit = func(current any) bool {
		switch typed := current.(type) {
		case string:
			return StoryProposalMatchesTextV1(proposal, typed)
		case []any:
			for _, item := range typed {
				if visit(item) {
					return true
				}
			}
		case map[string]any:
			var subjectValues, objectValues, stateValues []string
			for key, item := range typed {
				text, scalar := item.(string)
				if !scalar {
					continue
				}
				lower := strings.ToLower(strings.TrimSpace(key))
				switch {
				case proposalSubjectKeyV1(lower):
					subjectValues = append(subjectValues, text)
				case proposalObjectKeyV1(lower):
					objectValues = append(objectValues, text)
				case lower == "status" || lower == "authority" || lower == "certainty" || lower == "state":
					stateValues = append(stateValues, text)
				}
			}
			state := strings.Join(stateValues, " ")
			if !proposalNegativeOrUnknownV1(state) && proposalExactTermMatchV1(proposal.SubjectTerms, subjectValues) {
				for _, object := range objectValues {
					if !proposalNegativeOrUnknownV1(object) && proposalExactTermMatchV1(proposal.ObjectTerms, []string{object}) {
						return true
					}
				}
			}
			for _, item := range typed {
				if visit(item) {
					return true
				}
			}
		}
		return false
	}
	return visit(value)
}

func storyProposalWithDerivedTermsV1(proposal StoryProposalV1) StoryProposalV1 {
	if len(proposal.SubjectTerms) != 0 && len(proposal.PredicateTerms) != 0 && len(proposal.ObjectTerms) != 0 {
		return proposal
	}
	subjects, predicates, objects, err := deriveStoryProposalClaimV1(proposal.Statement)
	if err == nil {
		proposal.SubjectTerms, proposal.PredicateTerms, proposal.ObjectTerms = subjects, predicates, objects
	}
	return proposal
}

func DeriveStoryProposalClaimTermsV1(statement string) ([]string, []string, error) {
	subjects, _, objects, err := deriveStoryProposalClaimV1(statement)
	return subjects, objects, err
}

func deriveStoryProposalClaimV1(statement string) ([]string, []string, []string, error) {
	statement = strings.TrimSpace(statement)
	bestIndex, bestVerb := -1, ""
	for _, verb := range []string{"具有", "具备", "拥有", "属于", "能够", "可以", "是"} {
		if index := strings.Index(statement, verb); index >= 1 && (bestIndex < 0 || index < bestIndex) {
			bestIndex, bestVerb = index, verb
		}
	}
	if bestIndex < 1 {
		return nil, nil, nil, fmt.Errorf("statement requires an explicit subject-predicate-object claim")
	}
	subject := strings.Trim(strings.TrimSpace(statement[:bestIndex]), "，。！？；：,.!?;:\"'()（）[]【】")
	object := strings.Trim(strings.TrimSpace(statement[bestIndex+len(bestVerb):]), "，。！？；：,.!?;:\"'()（）[]【】")
	if subject == "" || object == "" {
		return nil, nil, nil, fmt.Errorf("statement requires non-empty subject and object")
	}
	objectTerms := []string{object}
	for _, suffix := range []string{"属性", "能力", "特性", "技能", "身份"} {
		if stem := strings.TrimSpace(strings.TrimSuffix(object, suffix)); stem != object && stem != "" {
			objectTerms = append(objectTerms, stem)
		}
	}
	predicateTerms := []string{bestVerb}
	switch bestVerb {
	case "具有", "具备", "拥有":
		predicateTerms = append(predicateTerms, "具有", "具备", "拥有", "是", "为")
	case "属于", "是":
		predicateTerms = append(predicateTerms, "属于", "是", "为")
	case "能够", "可以":
		predicateTerms = append(predicateTerms, "能够", "可以", "会")
	}
	return compactStoryProposalStrings([]string{subject}), compactStoryProposalStrings(predicateTerms), compactStoryProposalStrings(objectTerms), nil
}

func storyProposalMatchesClaimUnitV1(proposal StoryProposalV1, value string) bool {
	normalized := NormalizeProposalClaimTextV1(value)
	fullObject := ""
	if len(proposal.ObjectTerms) > 0 {
		fullObject = NormalizeProposalClaimTextV1(proposal.ObjectTerms[0])
	}
	for _, subject := range proposal.SubjectTerms {
		subject = NormalizeProposalClaimTextV1(subject)
		if subject == "" || fullObject == "" {
			continue
		}
		for _, predicate := range proposal.PredicateTerms {
			predicate = NormalizeProposalClaimTextV1(predicate)
			if predicate != "" && strings.Contains(normalized, subject+predicate+fullObject) {
				return true
			}
		}
		for _, suffix := range []string{"属性", "能力", "特性", "技能", "身份"} {
			suffix = NormalizeProposalClaimTextV1(suffix)
			stem := strings.TrimSuffix(fullObject, suffix)
			if stem != "" && stem != fullObject &&
				(strings.Contains(normalized, subject+"的"+suffix+"是"+stem) || strings.Contains(normalized, subject+"的"+suffix+"为"+stem)) {
				return true
			}
		}
	}
	return false
}

func proposalExactTermMatchV1(terms, values []string) bool {
	for _, value := range values {
		value = NormalizeProposalClaimTextV1(value)
		for _, term := range terms {
			if value != "" && value == NormalizeProposalClaimTextV1(term) {
				return true
			}
		}
	}
	return false
}

func proposalNegativeOrUnknownV1(value string) bool {
	lower := strings.ToLower(strings.TrimSpace(value))
	for _, marker := range []string{"不得", "不能", "不具有", "不具备", "没有", "无火", "尚未", "未批准", "未确定", "未建立", "未知", "候选", "提案", "proposal", "pending", "tbd", "undecided", "not established", "not accepted", "not approved"} {
		if strings.Contains(lower, marker) {
			return true
		}
	}
	return false
}

func proposalSubjectKeyV1(key string) bool {
	for _, marker := range []string{"character", "subject", "entity", "actor", "name", "protagonist", "speaker"} {
		if key == marker || strings.HasSuffix(key, "_"+marker) {
			return true
		}
	}
	return false
}

func proposalObjectKeyV1(key string) bool {
	for _, marker := range []string{"attribute", "property", "object", "value", "ability", "trait", "power", "capability"} {
		if key == marker || strings.HasSuffix(key, "_"+marker) {
			return true
		}
	}
	return false
}

func storyProposalDigest(value any) string {
	raw, _ := json.Marshal(value)
	sum := sha256.Sum256(raw)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func compactStoryProposalStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	out := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, duplicate := seen[value]; duplicate {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}
	return out
}
