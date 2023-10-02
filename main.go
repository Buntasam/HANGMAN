package main
  
import "fmt"

func main() {
	const (
		ligne int = 8
		colonne int = 20
	)

	var attempmax = 10
	var secretword = "thesecretword"
	
	var affichage [ligne][colonne]string
	for i := 0; i < ligne; i++ {
		for j := 0; j < colonne; j++ {
			fmt.Print(affichage[i][j])
		}
		fmt.Print("\n")
	}
}