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
	fmt.Println("Remaining attempts:", g.attemptsLeft)
	fmt.Println("Already used letters:", g.guesses)
	fmt.Print("Type a letter : ")
}

func (g *Game) verif() {

	g.Phrases()

	var text string
	fmt.Scanln(&text)
	fmt.Println()

	for i := 0; i < len(g.guesses); i++ {
		if len(text) != 1 || text == g.guesses[i] {
				fmt.Println("\n\n\n\n\n\n")
				fmt.Println("Not the expected answer")
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
		fmt.Println("\n\n\n\n")

		os.Exit(0)
	}

}


