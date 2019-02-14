package main

type ORDER_TYPE_ENUM uint8

const (
	order_hold        ORDER_TYPE_ENUM = iota
	order_move        ORDER_TYPE_ENUM = iota
	order_attack      ORDER_TYPE_ENUM = iota
	order_attack_move ORDER_TYPE_ENUM = iota
	order_build       ORDER_TYPE_ENUM = iota // maybe merge build and repair?
	order_construct   ORDER_TYPE_ENUM = iota
)

type order struct {
	orderType  ORDER_TYPE_ENUM
	x, y       int
	targetPawn *pawn

	buildingHasBeenPlaced bool // for build orders
	buildingToConstruct   *pawn
	constructingQueue     []*pawn // for units
}

func (clone *order) cloneFrom(o *order) {
	clone.orderType = o.orderType
	clone.x, clone.y = o.x, o.y
	clone.targetPawn = o.targetPawn
	clone.buildingHasBeenPlaced = o.buildingHasBeenPlaced
	clone.buildingToConstruct = o.buildingToConstruct
}

func (o *order) canBeDrawnAsLine() bool {
	return o.orderType != order_hold && o.orderType != order_construct
}
