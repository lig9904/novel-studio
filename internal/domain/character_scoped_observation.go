package domain

import "fmt"

const CharacterScopedObservationPolicyV1 = "character-scoped-observation-channel:v1"

func HasCharacterScopedObservationPolicyV1(sources []string) bool {
	return physicalContainsRefV2(sources, CharacterScopedObservationPolicyV1)
}

func validateCharacterScopedObservationStimulusV1(stimulus WorldStimulusPacket) error {
	policy := HasCharacterScopedObservationPolicyV1(stimulus.Sources)
	if policy && (stimulus.Version != WorldStimulusPacketV2Version || !HasCharacterOperationalAvailabilityPolicyV1(stimulus.Sources) || !HasCharacterSelfExperiencePolicyV2(stimulus.Sources)) {
		return fmt.Errorf("scoped observation channels require explicit v2 operational/self-experience policies")
	}
	if stimulus.PhysicalState == nil {
		return nil
	}
	mechanisms := map[string]CodexMechanism{}
	for _, mechanism := range stimulus.Mechanisms {
		mechanisms[mechanism.ID] = mechanism
	}
	for _, actor := range stimulus.PhysicalState.Actors {
		for _, holding := range actor.Resources {
			if len(holding.ObservationChannels) == 0 {
				continue
			}
			if !policy {
				return fmt.Errorf("actor-scoped observation channels require their frozen source policy")
			}
			for _, channel := range holding.ObservationChannels {
				mechanism, ok := mechanisms[channel.MechanismRef]
				if !ok || !characterResourceObservationMechanismAllowedV1(holding.Access, holding.Permissions, holding.ObservationChannels, mechanism) {
					return fmt.Errorf("actor-scoped observation channel references an unavailable or visibility-mismatched world mechanism")
				}
			}
		}
	}
	return nil
}

func validateCharacterScopedObservationPacketV1(observation CharacterObservationPacket) error {
	policy := HasCharacterScopedObservationPolicyV1(observation.Sources)
	for _, view := range observation.ResourceViews {
		if len(view.Permissions) == 0 && len(view.ObservationChannels) == 0 {
			continue
		}
		if !policy {
			return fmt.Errorf("resource permissions and observation channels require their frozen source policy")
		}
		if err := validateCharacterResourcePermissionsV1(view.Access, view.Permissions, view.ObservationChannels); err != nil {
			return err
		}
	}
	return nil
}
