package main

import (
	ge "game/gameEngine"
	"fmt"
)

func main() {
	var g ge.Game
	g.Init_game(8,20)
	for i := 0; i < g.Ligne; i++ {
		fmt.Println(g.Tableau[i])
	}
	g.Display()
}
