package main

type pawnWeaponInformation struct {
	attackRadius, attackDelay, attackEnergyCost int
	attacksLand, attacksAir bool // these are not mutually excluding
	hitscan *WeaponHitscan // TODO: non-hitscan (projectile) weapons
	canBeFiredOnMove bool
	nextTurnToFire int
}
