package agents

import "github.com/chenhongyang/novel-studio/internal/domain"

const characterSoftEventReadinessPromptV1 = `
软事件结果必须明确分类为OCCURRED、REJECTED_WITH_CONSEQUENCE、SUPERSEDED_BY_ACTUAL_CHOICE、DEFERRED或HARD_CONTRACT_UNSATISFIED。角色合法拒绝软纲本身不算完成：只有逐字引用该角色实际decision_reason，并逐字引用同一提案已裁决的immediate_result或state_after，且同周期真实时间或状态已经推进，才能判REJECTED_WITH_CONSEQUENCE。实际选择替代软纲时用SUPERSEDED_BY_ACTUAL_CHOICE，同样必须有角色理由和真实后果；不得继续强迫原软事件。没有形成实际后果用DEFERRED并继续。硬合同已被实际证明不可能才用HARD_CONTRACT_UNSATISFIED；角色自主不能覆盖硬合同。Host会核验actor_ref、proposal_ref、character_reason、world_consequence及同周期证据，禁止改写或概括这些字段。`

func characterActivationV3SoftEventReadinessPolicies() []string {
	return append(characterActivationV3ScopedObservationPolicies(), domain.CharacterSoftEventReadinessPolicyV1)
}

func characterActivationProtocolV3SoftEventReadinessDigest() string {
	digest, err := domain.DeterministicPlanningHash(struct {
		Base, Policy, ReadinessPrompt, ModelViewPolicy, SchemaPolicy string
		Schema                                                       map[string]any
	}{characterActivationProtocolV3ScopedObservationDigest(), domain.CharacterSoftEventReadinessPolicyV1, characterSoftEventReadinessPromptV1,
		domain.CharacterReadinessModelViewPolicyV2, domain.CharacterReadinessGroupedSchemaPolicyV2, domain.CharacterReadinessGroupedVerdictSchemaV2()})
	if err != nil {
		return ""
	}
	return "sha256:" + digest
}
