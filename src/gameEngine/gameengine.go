package gameEngine

import (
	"fmt"
)

type Game struct {
	Ligne int
	Colonne int
	Tableau [][]string
	stage int
	random_number int
	mot_secret []string
}



func (g *Game) Init_game(Ligne int, Colonne int) {
	g.Ligne = Ligne
	g.Colonne = Colonne
	g.stage = 0
	for i := 0; i < g.Ligne; i++ {
		var s []string
		g.Tableau = append(g.Tableau, s)
		for j := 0; j < g.Colonne; j++ {
			g.Tableau[i] = append(g.Tableau[i], " ")
		}
	}
	g.random_number = 0
}

func (g *Game) Running_game() {
	fmt.Println("Bienvenue sur notre pendu !")
	fmt.Println("Voulez-vous jouer ? yes / no")
	fmt.Println()
	g.Main_Menu()
}
