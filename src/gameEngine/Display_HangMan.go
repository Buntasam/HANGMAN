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
	g.mot = lines[g.Random_number]
	g.Tableau_rune = []rune(g.mot)
	fmt.Println()

	for i := 0; i < len(lines[g.Random_number]); i++ {
		g.Mot_secret = append(g.Mot_secret, "_")
	}
	
	fmt.Println()
}

func (g *Game) getGuess(lettre string) string {
	g.guesses = append(g.guesses, lettre)
	return g.guess
}

func (g *Game) Phrases() {
	fmt.Println()
	fmt.Println(g.Mot_secret)

	fmt.Print("\n\n")
	fmt.Println()
	fmt.Println("Tentatives restantes:", g.attemptsLeft)
	fmt.Println("Already used letters:", g.guesses)
	fmt.Print("Entrez une letttre : ")
}

func (g *Game) verif() {

	g.Phrases()

	var text string
	fmt.Scanln(&text)
	fmt.Println()

	for i := 0; i < len(g.guesses); i++ {
		if len(text) != 1 || text == g.guesses[i] {
				fmt.Println("\n\n\n\n")
				fmt.Println("Not good")
				g.verif()
				g.Display()
		}	
	}
	
	
	g.getGuess(text)
	g.perdu = true
	for j := 0; j < len(g.Tableau_rune); j++ {
		if string(text) == string(g.Tableau_rune[j]) {
			g.Mot_secret[j] = text
			g.perdu = false
		}
	}

	if g.perdu {
		g.Stage++
		g.attemptsLeft--
	}
}

func (g* Game) Win() {
	g.win = true
	for i := 0; i < len(g.Mot_secret); i++ {
		if g.Mot_secret[i] == "_" {
			g.win = false
		} 
	}
	if g.win {
		fmt.Print("\n\n\n\n\n\n\n\n")
		fmt.Println(" __     ______  _    _  __          _______ _   _ ")
		fmt.Println(" \\ \\   / / __ \\| |  | | \\ \\        / /_   _| \\ | |")
		fmt.Println("  \\ \\_/ / |  | | |  | |  \\ \\  /\\  / /  | | |  \\| |")
		fmt.Println("   \\   /| |  | | |  | |   \\ \\/  \\/ /   | | | . ` |")
		fmt.Println("    | | | |__| | |__| |    \\  /\\  /   _| |_| |\\  |")
		fmt.Println("    |_|  \\____/ \\____/      \\/  \\/   |_____|_| \\_|")
		fmt.Println("\n\n\n")

		os.Exit(0)
	}

}


func (g *Game) Display() {

	switch g.Stage {
	case 0: //vide
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}
		g.verif()
		g.Win()
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
		g.Win()
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
		g.Win()
		g.Display()

	case 3: //barre du haut
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		for i := 4; i < 12; i++ {
			g.Tableau[0][i] = "-"
		}
		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()

	case 4: //soutien
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][6] = "/"
		g.Tableau[2][5] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()

	case 5: //corde
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[1][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()

	case 6: //tete
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[2][11] = "O"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()

	case 7: // corps
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][11] = "|"
		g.Tableau[4][11] = "|"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()


	case 8: // bras droit
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()
		

	case 9: // bras gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[3][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()


	case 10: // jambe gauche
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][10] = "/"

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()
		

	case 11: // jambe droite
		fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		g.Tableau[5][12] = `\`

		for i := 0; i < len(g.Tableau); i++ {
			fmt.Println(g.Tableau[i])
		}

		g.verif()
		g.Win()
		g.Display()
		

	case 12: // game over
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
		fmt.Printf("Le mot était : %s", g.mot)
		fmt.Println()
		fmt.Println()
	}
}
