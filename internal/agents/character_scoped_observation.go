package agents

import (
	"github.com/chenhongyang/novel-studio/internal/domain"
	"github.com/chenhongyang/novel-studio/internal/tools"
)

const characterScopedObservationPromptV1 = `
角色限定观察通道：resource_views.permissions把observe/use/possess/control/claim分开；能观察不等于拥有、控制、认领、可进入或可使用。operational_observation只能选本人resource_view中允许observe的资源。resource_view若列observation_channels，本次观察mechanism_ref必须取其中一项；private通道只把该条安全label和机制ID提供给本人，不公开作者态机制、其他角色身份或隐藏结果。没有实际执行仍不会产生观察结果。`

const worldArbiterScopedObservationPromptV1 = `
角色限定观察裁决：公开机制可用于明确获observe权限的非所有者；private观察通道只授权绑定角色对绑定资源使用指定秘密机制，不授予use/possess/control/claim，也不向其他角色投放机制或观察结果。严格核对角色、resource_id、mechanism_ref、实际位置、工作区间和结果时点；不得因同场、共享资源或作者态秘密自动复制知识。`

func characterActivationV3ScopedObservationPolicies() []string {
	return append(characterActivationV3IncomingReadPolicies(), domain.CharacterScopedObservationPolicyV1)
}

func characterActivationProtocolV3ScopedObservationDigest() string {
	policies := append(characterActivationV3ScopedObservationPolicies(), domain.CharacterSourceRefPolicyV2, domain.CharacterSelfExperiencePolicyV2, domain.CharacterOperationalAvailabilityPolicyV1, domain.CharacterPassiveReceptionPolicyV2)
	submit := tools.NewSubmitCharacterDecisionTool(nil, domain.CharacterObservationPacket{Version: domain.CharacterObservationV2Version, Sources: policies})
	token, _ := domain.CharacterActivationCycleSourceToken("pg2_scoped_observation_schema", 1, 1, "sha256:0000000000000000000000000000000000000000000000000000000000000000", "")
	resolve := tools.NewResolveChapterWorldTool(nil, domain.WorldStimulusPacket{Version: domain.WorldStimulusPacketV2Version, PhysicalState: &domain.WorldPhysicalStateV2{}, StoryClock: &domain.StoryClockContext{}, Sources: append(policies, token)}, domain.CharacterAgentActivation{}, nil, "", nil, 1)
	digest, err := domain.DeterministicPlanningHash(struct {
		Base, Policy, CharacterPrompt, ArbiterPrompt string
		Submit, Resolve                              map[string]any
	}{characterActivationProtocolV3IncomingReadDigest(), domain.CharacterScopedObservationPolicyV1, characterScopedObservationPromptV1, worldArbiterScopedObservationPromptV1, submit.Schema(), resolve.Schema()})
	if err != nil {
		return ""
	}
	return "sha256:" + digest
}
