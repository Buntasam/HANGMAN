package gameEngine

import (
	"fmt"
	"math/rand"
)

type Game struct {
	Ligne         int
	Colonne       int
	Tableau       [][]string
	Stage         int
	Random_number int
	Mot_secret    []string
	Content       []byte
	Err           error
	Lines         []string
	S1            rand.Source
	R1            *rand.Rand
	Tableau_rune  []rune
	perdu         bool
	mot           string
	win           bool
}

func (g *Game) Init_game(Ligne int, Colonne int) {
	g.Ligne = Ligne
	g.Colonne = Colonne
	g.Stage = 0
	for i := 0; i < g.Ligne; i++ {
		var s []string
		g.Tableau = append(g.Tableau, s)
		for j := 0; j < g.Colonne; j++ {
			g.Tableau[i] = append(g.Tableau[i], " ")
		}
	}
	g.Random_number = 0
}

func (g *Game) Running_game() {
	fmt.Println("  _    _          _   _  _____ __  __          _   _ ")
	fmt.Println(" | |  | |   /\\   | \\ | |/ ____|  \\/  |   /\\   | \\ | |")
	fmt.Println(" | |__| |  /  \\  |  \\| | |  __| \\  / |  /  \\  |  \\| |")
	fmt.Println(" |  __  | / /\\ \\ | . ` | | |_ | |\\/| | / /\\ \\ | . ` |")
	fmt.Println(" | |  | |/ ____ \\| |\\  | |__| | |  | |/ ____ \\| |\\  |")
	fmt.Println(" |_|  |_/_/    \\_\\_| \\_|\\_____|_|  |_/_/    \\_\\_| \\_|")
	fmt.Println()
	fmt.Println("Bienvenue sur notre pendu !")
	fmt.Println("Voulez-vous jouer ? yes / no")
	fmt.Println()
	g.Main_Menu()
}
