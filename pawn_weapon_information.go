package main

type pawnWeaponInformation struct {
	attackRadius, attackDelay int
	attacksLand, attacksAir bool // these are not mutually excluding
	hitscan *WeaponHitscan // TODO: non-hitscan (projectile) weapons
}
