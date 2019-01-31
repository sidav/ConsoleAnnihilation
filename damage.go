package main

import "fmt"

func dealDamageToTarget(attacker *pawn, wpn *pawnWeaponInformation, target *pawn) {
	damageDealt := wpn.hitscan.baseDamage
	if target.isLight {
		damageDealt += wpn.hitscan.lightMod
	}
	if target.isHeavy {
		damageDealt += wpn.hitscan.heavyMod
	}
	target.hitpoints -= damageDealt
	log.appendMessage(fmt.Sprintf("%s pewpews at %s (%d damage)", attacker.name, target.name, damageDealt))
}
