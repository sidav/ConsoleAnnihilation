package main 

type squadMember struct {
	code string
}

func (sm *squadMember) getStaticData() *squadMemberInfo {
	return getSquadMemberStaticInfo(sm.code)
}
