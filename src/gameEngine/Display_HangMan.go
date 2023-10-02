package gameEngine

import "fmt"

func (g *Game) Display() {
	switch g.stage {
	case 0: //sol
		fmt.Print(g.Tableau)
	case 1: //poteau
		for i := 0; i < 9; i++ {
			g.Tableau[g.Ligne-1][i] = "-"
		}
		fmt.Print(g.Tableau)
	case 2: //barre du haut
		for i := 0; i < 8; i++ {
			g.Tableau[i][4] = "|"
		}
		fmt.Print(g.Tableau)

	case 3: //soutien

	case 4: //corde

	case 5: //tete

	case 6: //corps

	case 7: // bras gauche

	case 8: // bras droit

	case 9: // jambe gauche

	case 10: // jambe droite

	}

}
