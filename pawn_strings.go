package main 

import "fmt"

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
