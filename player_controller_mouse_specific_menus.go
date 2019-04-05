package main

import (
	"fmt"
	cmenu "github.com/sidav/golibrl/console_menu"
	"sort"
)

var pcm_lastKeyEquivalent string

func pcm_mouseOrderSelectMenu(p *pawn) string {
	orders := make([]string, 0)
	keyEquivalents := make([]string, 0)
	if p.canConstructBuildings() {
		orders = append(orders, "(B)uild")
		keyEquivalents = append(keyEquivalents, "b")
	}
	if p.canConstructUnits() {
		orders = append(orders, "(C)onstruct units")
		keyEquivalents = append(keyEquivalents, "c")
		keyEquivalents = append(keyEquivalents, "r")
		if p.repeatConstructionQueue {
			orders = append(orders, "(R)epeat queue: ENABLED")
		} else {
			orders = append(orders, "(R)epeat queue: DISABLED")
		}
		if p.faction.cursor.currentCursorMode == CURSOR_AMOVE {
			orders = append(orders, "Default order: attack-move")
			orders = append(orders, "(m): set to move")
			keyEquivalents = append(keyEquivalents, "m")
			keyEquivalents = append(keyEquivalents, "m")
		} else {
			orders = append(orders, "Default order: move")
			orders = append(orders, "(a): set to attack-move")
			keyEquivalents = append(keyEquivalents, "a")
			keyEquivalents = append(keyEquivalents, "a")
		}
	}
	if p.faction.cursor.currentCursorMode == CURSOR_AMOVE && p.canMove() {
		orders = append(orders, "(M)ove")
		keyEquivalents = append(keyEquivalents, "m")
	} else {
		if p.hasWeapons() {
			orders = append(orders, "(A)ttack-move")
			keyEquivalents = append(keyEquivalents, "a")
		}
	}
	index := cmenu.DrawSidebarMouseOnlyAsyncMenu("Orders for: "+p.name, p.faction.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_3, SIDEBAR_W, orders)
	if index == -1 {
		return "NONE"
	}
	return keyEquivalents[index]
}

func pcm_MouseOrderForMultiselectMenu(f *faction, selection *[]*pawn) string {
	orders := make([]string, 0)
	keyEquivalents := make([]string, 0)
	selectedUnitsCounter := make(map[string]int)
	for _, p := range *selection {
		selectedUnitsCounter[p.name]++
	}
	// sort the map because of dumbass Go developers thinking that they know your needs better than you do
	keys := make([]string, 0, len(selectedUnitsCounter))
	for k := range selectedUnitsCounter {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		orders = append(orders, fmt.Sprintf("%dx %s", selectedUnitsCounter[key], key))
		keyEquivalents = append(keyEquivalents, "NONE")
	}
	if f.cursor.currentCursorMode == CURSOR_AMOVE {
		orders = append(orders, "(M)ove")
		keyEquivalents = append(keyEquivalents, "m")
	} else {
		orders = append(orders, "(A)ttack-move")
		keyEquivalents = append(keyEquivalents, "a")
	}
	index := cmenu.DrawSidebarMouseOnlyAsyncMenu(fmt.Sprintf("ORDERS FOR %d UNITS", len(*selection)), f.getFactionColor(),
		SIDEBAR_X, SIDEBAR_FLOOR_3, SIDEBAR_W, orders)
	if index == -1 {
		return "NONE"
	}
	return keyEquivalents[index]
}
