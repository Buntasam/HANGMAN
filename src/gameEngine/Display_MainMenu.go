package gameEngine

import "fmt"



func (g *Game) Main_Menu() {
	var text string
	fmt.Scanln(&text)
	switch text {
	case "yes":
		fmt.Println()
		g.aléatoire()
		g.Display()
	default:
		fmt.Println("TRY AGAIN ")
		fmt.Println()
		g.Main_Menu()
	}

}