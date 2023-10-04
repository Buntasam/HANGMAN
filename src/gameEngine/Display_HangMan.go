package gameEngine

import "fmt"

func (g *Game) Display() {
	switch g.stage {
	case 0: //vide
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "1" {
			g.stage = 1
			g.Display()
		}

	case 1: //sol
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 9; i++ {
			g.Tableau[g.Ligne-1][i] = "-"
		}

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "2" {
			g.stage = 2
			g.Display()
		}

	case 2: //poteau
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 7; i++ {
			g.Tableau[i][4] = "|"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "3" {
			g.stage = 3
			g.Display()
		}

	case 3: //barre du haut
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 4; i < 12; i++ {
			g.Tableau[0][i] = "-"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "4" {
			g.stage = 4
			g.Display()
		}

	case 4: //soutien
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][6] = "/"
		g.Tableau[2][5] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "5" {
			g.stage = 5
			g.Display()
		}

	case 5: //corde
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "6" {
			g.stage = 6
			g.Display()
		}

	case 6: //tete
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[2][11] = "O"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "7" {
			g.stage = 7
			g.Display()
		}

	case 7: // corps
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][11] = "|"
		g.Tableau[4][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "8" {
			g.stage = 8
			g.Display()
		}

	case 8: // bras droit
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "9" {
			g.stage = 9
			g.Display()
		}

	case 9: // bras gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "10" {
			g.stage = 10
			g.Display()
		}

	case 10: // jambe gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "11" {
			g.stage = 11
			g.Display()
		}

	case 11: // jambe droite
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		fmt.Print("\n\n")
		var text string
		fmt.Scanln(&text)

		if text == "12" {
			g.stage = 12
			g.Display()
		}
	
	case 12: // jambe droite
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][12] = `\`
		fmt.Println("  ________                        ________                     ")
		fmt.Println(" /  _____/_____    _____   ____  ",`\`+"_____ ",`\`,"___  __ ___________ ")
		fmt.Println("/  ",`\`," ___"+`\`+"__",`\`," /    ",`\`+"_/ __",`\`,"   /  |  ",`\`," ",`\`+"/ // __",`\`+"_  __",`\`)
		fmt.Println(`\`,"  ",`\`+"_"+`\ `,`\`,"/ __ ",`\`,"|  Y Y  ",`\`,"  ___/  /    |    ",`\`,"/",`\`,"  ___/|  | ",`\`,"/")
		fmt.Println(" ",`\`,"______  (____  /__|_|  /",`\`,"___  > ",`\`,"_______  /",`\`,"_/  ",`\`,"___  >__|   ")
		fmt.Println(        \/     \/      \/     \/          \/          \/       )
		fmt.Println("\n")
	}
}
