package gameEngine

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (g *Game) aléatoire() {
	g.Content, g.Err = os.ReadFile("gameEngine/mot.txt")
	if g.Err != nil {
		fmt.Print(g.Err)
	}
	lines := strings.Split(string(g.Content), "\n")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	nb_mot := len(lines)
	g.Random_number = r1.Intn(nb_mot)
	fmt.Println(lines[g.Random_number])
	g.Tableau_rune = []rune(lines[g.Random_number])
	fmt.Println(g.Tableau_rune)
	fmt.Println()

	for i := 0; i < len(lines[g.Random_number]); i++ {
		g.Mot_secret = append(g.Mot_secret, "_")
	}

	fmt.Println(g.Mot_secret)
	fmt.Println()
}

func (g *Game) verif() {
	fmt.Println(g.Mot_secret)

	fmt.Print("\n\n")
	fmt.Print("Entrez une letttre : ")
	var text string
	fmt.Scanln(&text)
	fmt.Println()
	g.perdu = true
	for j := 0; j < len(g.Tableau_rune)-1; j++ {
		if string(text) == string(g.Tableau_rune[j]) {
			fmt.Println("vrai")
			g.Mot_secret[j] = text
			g.perdu = false
		}
	}
	if g.perdu {
		g.Stage++
	}
}

func (g *Game) Display() {

	switch g.Stage {
	case 0: //vide
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		g.verif()
		g.Display()

	case 1: //sol
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 9; i++ {
			g.Tableau[g.Ligne-1][i] = "-"
		}

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Display()

	case 2: //poteau
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 0; i < 7; i++ {
			g.Tableau[i][4] = "|"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		g.verif()
		g.Display()

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
			g.Stage = 4
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
			g.Stage = 5
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
			g.Stage = 6
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
			g.Stage = 7
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
			g.Stage = 8
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
			g.Stage = 9
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
			g.Stage = 10
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
			g.Stage = 11
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
			g.Stage = 12
			g.Display()
		}

	case 12: // jambe droite
		fmt.Print("\n\n")
		g.Tableau[5][12] = `\`
		fmt.Println("  ________                        ________                     ")
		fmt.Println(" /  _____/ ____    _____   ____  ", `\`+"_____ ", `\`+"___  __ ___________ ")
		fmt.Println("/  ", `\`, " ___", `\`+"__", `\`, " /    ", `\`+"_/ __", `\`, "  /   |  ", `\`+" ", `\`+"/ // __", `\`+"_  __", `\`)
		fmt.Println(`\`, "  ", `\`+"_"+`\ `, `\`+"/ __", `\`+"|  Y Y ", `\`, " ___/  /    |   ", `\`, "  /", `\`, "___/|  | "+`\`+"/")
		fmt.Println(` \` + "______  (____  /__|_|  /" + `\` + "___  > " + `\` + "_______  /" + `\` + "_/   " + `\` + "___ >__|   ")
		fmt.Println("       ", `\`+"/    ", `\`+"/      "+`\`+"/    ", `\`+"/         ", `\`+"/           "+`\`+"/       ")
		fmt.Println()
		fmt.Println()
	}
}
