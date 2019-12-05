package main

//contains squad member data.
type squadMemberStaticData struct {
	// movementDelay int
	name                     string
	maxHp                    int
	sightRadius, radarRadius int

	size            int
	takesWholeSquad bool

	movementInfo              *pawnMovementInformation
	weaponInfo                *pawnWeaponInformation
	res                       *pawnResourceInformation
	nanolatherInfo            *nanolatherInformation
	currentConstructionStatus *constructionInformation
}
