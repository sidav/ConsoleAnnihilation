package main

//contains squad member data.
type squadMemberInfo struct {
	// movementDelay int
	name            string
	maxHp           int
	size            int
	takesWholeSquad bool 

	movementInfo              *pawnMovementInformation
	weaponInfo                *pawnWeaponInformation
	res                       *pawnResourceInformation
	nanolatherInfo            *nanolatherInformation
	currentConstructionStatus *constructionInformation
}
