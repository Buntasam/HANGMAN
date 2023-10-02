package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const (
		ligne int = 8
		colonne int = 20
	)

	//var attempmax = 10
	var secretword = "thesecretword"
	var currentword []string

	for i := 0; i < len(secretword); i++ {
		currentword = append(currentword, "_")
	}


	var affichage [ligne][colonne]string
	for i := 0; i < ligne; i++ {
		for j := 0; j < colonne; j++ {
			fmt.Print(affichage[i][j])
		}
		fmt.Print("\n")
	}
	

	fmt.Print(currentword)
	fmt.Print("\n")
	fmt.Print("\n")

	for i := 0; i < len(currentword); i++ {
		fmt.Print(currentword[i], " ")
	}
	fmt.Print("\n")
	fmt.Print("\n")

	currentword[3] = "I"

	for i := 0; i < len(currentword); i++ {
		fmt.Print(currentword[i], " ")
	}
	fmt.Print("\n")
	fmt.Print("\n")


	file, err := ioutil.ReadFile("mot.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	text := string(file)
	fmt.Println(text)
}

