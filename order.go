package main

type ORDER_TYPE_ENUM uint8

const (
	order_hold        ORDER_TYPE_ENUM = iota
	order_move        ORDER_TYPE_ENUM = iota
	order_attack      ORDER_TYPE_ENUM = iota
	order_attack_move ORDER_TYPE_ENUM = iota
	order_build       ORDER_TYPE_ENUM = iota // maybe merge build and repair?
)

type order struct {
	orderType             ORDER_TYPE_ENUM
	x, y                  int
	targetUnit            *unit
	targetBuilding        *pawn
	buildingHasBeenPlaced bool // for build orders
}
