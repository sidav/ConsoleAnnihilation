package main 

type squad struct {
	members []*squadMember
}

func (s *squad) getSquadName() string {
	// var smdata *squadMemberInfo
	smdata := s.members[0].getStaticData()
	return smdata.name
	if smdata.takesWholeSquad {
		return smdata.name
	}
	return "Strike team"
}

func (s *squad) getSquadMovementInfo() *pawnMovementInformation {
	movesOnLand := true 
	movesOnSea := true 
	movementDelay := 0 
	if len(s.members) > 0 {
		for i := range s.members {
			smInfo := getSquadMemberStaticInfo(s.members[i].code) 
			movInfo := smInfo.movementInfo
			if movInfo == nil {
				return nil 
			}
			movesOnLand = movesOnLand && movInfo.movesOnLand
			movesOnSea = movesOnSea && movInfo.movesOnSea
			if movInfo.ticksForMoveSingleCell > movementDelay {
				movementDelay = movInfo.ticksForMoveSingleCell
			}
		}
		return &pawnMovementInformation{ticksForMoveSingleCell: movementDelay, movesOnLand: movesOnLand, movesOnSea: movesOnSea}
	}
	return nil 
}

func (s *squad) getSquadSightAndRadarRadius() (int, int) {
	sr := 0
	rr := 0 
	for i := range s.members {
		static := getSquadMemberStaticInfo(s.members[i].code)
		if sr < static.sightRadius {
			sr = static.sightRadius
		}
		if rr < static.radarRadius {
			rr = static.radarRadius
		}
	}
	return sr, rr 
}

// form nanolather information 
func (s *squad) getSquadNanolatherInfo() *nanolatherInformation {
	// TODO: merge available allowed buildings lists for all nanolatherInfos. 
	for _, member := range s.members {
		static := getSquadMemberStaticInfo(member.code)
		if static.nanolatherInfo != nil {
			return static.nanolatherInfo
		}
	}
	return nil 
}

func (s *squad) getSquadIncome() *pawnIncomeInformation {
	var resultIncome *pawnIncomeInformation = nil 
	for _, sm := range s.members {
		static := getSquadMemberStaticInfo(sm.code)
		if static.income != nil {
			if resultIncome == nil {
				resultIncome = &pawnIncomeInformation{}
			}
			resultIncome.metalIncome += static.income.metalIncome
			resultIncome.energyIncome += static.income.energyIncome
			resultIncome.metalStorage += static.income.metalStorage
			resultIncome.energyStorage += static.income.energyStorage
			resultIncome.energyDrain += static.income.energyDrain
		}
	}
	return resultIncome
}
