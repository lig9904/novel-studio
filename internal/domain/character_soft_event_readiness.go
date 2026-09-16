package domain

const CharacterSoftEventReadinessPolicyV1 = "character-soft-event-readiness:v1"

func HasCharacterSoftEventReadinessPolicyV1(sources []string) bool {
	return physicalContainsRefV2(sources, CharacterSoftEventReadinessPolicyV1)
}
