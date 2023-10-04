package gameEngine



type Game struct {
	Ligne int
	Colonne int
	Tableau [][]string
	stage int
}



func (g *Game) Init_game(Ligne int, Colonne int) {
	g.Ligne = Ligne
	g.Colonne = Colonne
	g.stage = 12
	for i := 0; i < g.Ligne; i++ {
		var s []string
		g.Tableau = append(g.Tableau, s)
		for j := 0; j < g.Colonne; j++ {
			g.Tableau[i] = append(g.Tableau[i], " ")
		}
	}
}

