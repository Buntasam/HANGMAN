package main

import "strings"

type armature struct {
	Affichage [][]string

}

func (a *armature) Init() {
	var tab []string
	for i := 0; i < 8; i++ {
		a.Affichage[i] = append(a.Affichage[i], tab)
		for j := 0; j < 20; j++ {
			a.Affichage[i][j] = append(a.Affichage[i][j], "")
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 20; j++ {
			a.Affichage[i][j] = "0"
		}
	}
}