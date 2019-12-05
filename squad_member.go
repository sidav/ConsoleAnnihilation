package main 

type squadMember struct {
	code string
}

func (sm *squadMember) getStaticData() *squadMemberStaticData {
	return getSquadMemberStaticInfo(sm.code)
}
