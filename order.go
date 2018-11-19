package main

type ORDER_TYPE_ENUM uint8

const (
	order_hold ORDER_TYPE_ENUM = iota
	order_move ORDER_TYPE_ENUM = iota
	order_attack ORDER_TYPE_ENUM = iota
	order_attack_move ORDER_TYPE_ENUM = iota
)

type order struct {
	order_type ORDER_TYPE_ENUM
	x, y int
	target unit
}

