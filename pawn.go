package main

import (
	"fmt"
	geometry "github.com/sidav/golibrl/geometry"
)

type pawn struct {
	// pawn is a building or a unit squad - anything that can be commanded.
	squadInfo    *squad
	buildingInfo *building

	faction                   *faction
	x, y                      int
	order                     *order
	spending                  *pawnSpendings
	currentConstructionStatus *constructionInformation
	weapons                   []*pawnWeaponInformation
	nextTickToAct             int
	isCommander               bool

	repeatConstructionQueue bool // for factories
	// armor info:
	hitpoints        int
	isLight, isHeavy bool // these are not mutually excluding lol. Trust me, I'm a programmer
	regenPeriod      int  // if != 0 then the pawn will regen 1 hp each Nth tick.
}

func (p *pawn) hasWeapons() bool {
	return len(p.weapons) > 0
}

func (p *pawn) getMovementInfo() *pawnMovementInformation {
	if p.isSquad() { // TODO: remove isUnit
		return &pawnMovementInformation{ticksForMoveSingleCell: 10, movesOnLand: true}
	}
	return nil
}

func (p *pawn) getMaxRadiusToFire() int {
	max := 0
	for _, weap := range p.weapons {
		if weap.attackRadius > max {
			max = weap.attackRadius
		}
	}
	return max
}

func (p *pawn) isAlive() bool {
	if p.isSquad() {
		return len(p.squadInfo.members) > 0
	}
	return p.hitpoints > 0
}

func (p *pawn) isSquad() bool {
	return p.squadInfo != nil
}

func (p *pawn) isBuilding() bool {
	return p.buildingInfo != nil
}

func (p *pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) getSize() (int, int) {
	if p.isBuilding() {
		bs := p.buildingInfo.getBuildingStaticInfo()
		return bs.w, bs.h
	}
	return 1, 1
}

func (p *pawn) getName() string {
	if p.isSquad() {
		return p.squadInfo.getSquadName()
	}
	return p.buildingInfo.getName()
}

func (p *pawn) setOrder(o *order) {
	p.order = o
	if p.faction.playerControlled {
		log.appendMessage(fmt.Sprintf("%s order for %d, %d confirmed!", p.getCurrentOrderImperative(), o.x, o.y))
	}
}

func (p *pawn) getNanolatherInfo() *nanolatherInformation {
	if p.isSquad() {
		return p.squadInfo.getSquadNanolatherInfo()
	}
	return p.buildingInfo.nanolatherInfo
}

func (p1 *pawn) isInDistanceFromPawn(p2 *pawn, r int) bool {
	p1_w, p1_h := p1.getSize()
	p2_w, p2_h := p2.getSize()
	if p1.isBuilding() {
		if p2.isBuilding() {
			return geometry.AreRectsInRange(p1.x, p1.y, p1_w, p1_h, p2.x, p2.y, p2_w, p2_h, r)
		}
		return geometry.AreCoordsInRangeFromRect(p2.x, p2.y, p1.x, p1.y, p1_w, p1_h, r)
	} else {
		if p2.buildingInfo != nil {
			return geometry.AreCoordsInRangeFromRect(p1.x, p1.y, p2.x, p2.y, p2_w, p2_h, r)
		}
		return geometry.AreCoordsInRange(p1.x, p1.y, p2.x, p2.y, r)
	}
}

func (p *pawn) isOccupyingCoords(x, y int) bool {
	if p.isBuilding() {
		b_w, b_h := p.getSize()
		return geometry.AreCoordsInRect(x, y, p.x, p.y, b_w, b_h)
	} else {
		return x == p.x && y == p.y
	}
}

func (p *pawn) IsCloseupToCoords(x, y, dist int) bool { // distance to any of pawn's cells check.
	if p.isBuilding() {
		b_w, b_h := p.getSize()
		return !geometry.AreCoordsInRect(x, y, p.x, p.y, b_w, b_h) &&
			geometry.AreCoordsInRect(x, y, p.x-dist, p.y-dist, b_w+2*dist, b_h+2*dist)
	} else {
		return x != p.x && y != p.y && geometry.AreCoordsInRect(x, y, p.x-dist, p.y-dist, 2*dist+1, 2*dist+1)
	}
}

func (p *pawn) getCenter() (int, int) {
	if p.isSquad() {
		return p.x, p.y
	} else {
		b_w, b_h := p.getSize()
		return p.x + b_w/2, p.y + b_h/2
	}
}

func (p *pawn) getSightAndRadarRadius() (int, int) {
	if p.isSquad() {
		return p.squadInfo.getSquadSightAndRadarRadius()
	}
	si := p.buildingInfo.getBuildingStaticInfo()
	return si.sightRadius, si.radarRadius
}

func (p *pawn) getIncomeData() *pawnIncomeInformation {
	if p.isSquad() {
		return p.squadInfo.getSquadIncome()
	}
	return p.buildingInfo.getBuildingStaticInfo().defaultIncomeData
}

func (p *pawn) getMaxHitpoints() int {
	if p.isSquad() {
		return 10 // TODO: replace with staffing for squads 
	}
	return p.buildingInfo.getBuildingStaticInfo().maxHitpoints
}

func (p *pawn) getArmorDescriptionString() string {
	armorInfo := fmt.Sprintf("Armor %d / %d", p.hitpoints, p.getMaxHitpoints())
	if p.isLight {
		armorInfo += ", light"
	}
	if p.isHeavy {
		armorInfo += ", heavy"
	}
	if p.regenPeriod > 0 {
		regenPerTurnMult10 := 100 / (p.regenPeriod)
		armorInfo += fmt.Sprintf(", regen %d.%d", regenPerTurnMult10/10, regenPerTurnMult10%10)
	}
	return armorInfo
}

func (p *pawn) getCurrentOrderDescription() string {
	if p.currentConstructionStatus != nil {
		constr := p.currentConstructionStatus
		return fmt.Sprintf("UNDER CONSTRUCTION: %d%%", constr.getCompletionPercent())
	}
	if p.order == nil {
		return "STANDBY"
	}
	switch p.order.orderType {
	case order_hold:
		return "STANDBY"
	case order_move:
		return "MOVING"
	case order_attack:
		return "ASSAULTING"
	case order_attack_move:
		return "MOVING WHILE ENGAGING"
	case order_build:
		return fmt.Sprintf("NANOLATHING (%d%% ready)",
			p.order.buildingToConstruct.currentConstructionStatus.getCompletionPercent())
	case order_construct:
		if len(p.order.constructingQueue) > 0 {
			return fmt.Sprintf("CONSTRUCTING: %s (%d%% ready)", p.order.currentPawnUnderConstruction.getName(),
				p.order.currentPawnUnderConstruction.currentConstructionStatus.getCompletionPercent())
		} else {
			return "FINISHING CONSTRUCTION"
		}
	default:
		return "DOING SOMETHING"
	}
}

func (p *pawn) getCurrentOrderImperative() string {
	if p.order == nil {
		return "STAND BY"
	}
	switch p.order.orderType {
	case order_hold:
		return "STAND BY"
	case order_move:
		return "MOVE"
	case order_attack:
		return "ASSAULT"
	case order_attack_move:
		return "ATTACK MOVE"
	case order_build:
		return fmt.Sprintf("BUILD")
	case order_construct:
		return "CONSTRUCT"
	default:
		return "BLAH BLAH BLAH"
	}
}
