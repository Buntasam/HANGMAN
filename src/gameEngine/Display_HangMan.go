package gameEngine

import "fmt"

func (g *Game) Display() {
	switch g.stage {
	case 0: //sol
		fmt.Print(g.Tableau)
	case 1: //poteau
		
		fmt.Print(g.Tableau)
	case 2: //barre du haut

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