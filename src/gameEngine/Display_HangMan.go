package gameEngine

import "fmt"

func (g *Game) Display() {
	switch g.stage {
	case 0: //sol
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		var text string
		fmt.Scanln(&text)

		if text == "1" {
			g.stage = 1
			g.Display()
		}
	case 1: //poteau
		for i := 0; i < 9; i++ {
			g.Tableau[g.Ligne-1][i] = "-"
		}
		
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		var text string
		fmt.Scanln(&text)

		if text == "2" {
			g.stage = 2
			g.Display()
		}
	case 2: //barre du haut
		for i := 0; i < 7; i++ {
			g.Tableau[i][4] = "|"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

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
