package main

import (
	ge "game/gameEngine"
)

func main() {
	var g ge.Game
	g.Init_game(8,20)
	g.Display()
}
