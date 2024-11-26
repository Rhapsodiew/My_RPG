package gamecustom

import (
	"bufio"
	"fmt"
	"os"
)

func Title_screen() {
	fmt.Println("\033[33;1m\033[33;4m========== THE HYRULE CASTLE ==========\033[0m")
	fmt.Println()
	fmt.Println("\033[97;1mN. 	Nouvelle partie\033[0m")
	Option_mod()
	fmt.Println("\033[97;1mQ.	Quitter\033[0m")
	fmt.Println()
}

func Chose_difficulty_floor() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[33;1m-> \033[97mChoissisez une difficulté :\033[0m")
	fmt.Println()
	fmt.Println("\033[97;1m1. 	Normal (DEFAULT)\033[0m")
	fmt.Println("\033[97;1m2. 	Difficile\033[0m")
	fmt.Println("\033[97;1m3. 	Extreme\033[0m")
	fmt.Println()
	scanner.Scan()
	difficulty := scanner.Text()
	fmt.Println("\033[33;1m-> \033[97mChoissisez le nombre d'étage du chateau :\033[0m ")
	fmt.Println()
	fmt.Println("\033[97;1m1.     10 (DEFAULT)\033[0m")
	fmt.Println("\033[97;1m2.     20\033[0m")
	fmt.Println("\033[97;1m3.     50\033[0m")
	scanner.Scan()
	etage := scanner.Text()
	var chosen_etage int
	var chosen_difficulty int
	if difficulty != "1" && difficulty != "2" && difficulty != "3" {
		chosen_difficulty = 1
	} 
	if etage != "1" && etage != "2" && etage != "3" {
		chosen_etage = 1
	}
	if difficulty == "1" {
		chosen_difficulty = 1
	} else if difficulty == "2" {
		chosen_difficulty = 2
	} else if difficulty == "3" {
		chosen_difficulty = 3
	}
	if etage == "1" {
		chosen_etage = 10
	} else if etage == "2" {
		chosen_etage = 20
	} else if etage == "3" {
		chosen_etage = 50
	}
	return chosen_difficulty, chosen_etage
}

func Check_start(start string) (string, int, int) {
	if start == "N" || start == "Nouvelle partie" {
		difficulty, etage := Chose_difficulty_floor()
		return "N", difficulty, etage
	} else if start == "Q" || start == "Quit" {
		return "Q", 0, 0
	}
	return "o",0,0
}
