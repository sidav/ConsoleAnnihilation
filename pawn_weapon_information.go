package main

import "fmt"

type pawnWeaponInformation struct {
	attackRadius, attackDelay, attackEnergyCost int
	attacksLand, attacksAir bool // these are not mutually excluding
	hitscan *WeaponHitscan // TODO: non-hitscan (projectile) weapons
	canBeFiredOnMove bool
	nextTurnToFire int
}

func (wpn *pawnWeaponInformation) getDescriptionString() string{
	desc := fmt.Sprintf("Attack radius %d, delay %d.%d, ", wpn.attackRadius, wpn.attackDelay/10, wpn.attackDelay % 10)
	if wpn.hitscan != nil {
		desc += fmt.Sprintf("damage %d", wpn.hitscan.baseDamage)
		if wpn.hitscan.lightMod != 0 {
			desc += fmt.Sprintf(" (vs light %d)", wpn.hitscan.baseDamage + wpn.hitscan.lightMod)
		}
		if wpn.hitscan.heavyMod != 0 {
			desc += fmt.Sprintf(" (vs heavy %d)", wpn.hitscan.baseDamage + wpn.hitscan.heavyMod)
		}
		if wpn.attackEnergyCost != 0 {
			desc += fmt.Sprintf(", %d energy per shot", wpn.attackEnergyCost)
		}
	}
	return desc 
}
