package main

import (
	"bufio"
	"fmt"
	"mods/structure"
	"os"
)

func tutoriel(tuto string, p1 structure.Character, enemies []structure.Character, etage int, money int) (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	if tuto == "y" || tuto == "Y" {
		fmt.Println()
		fmt.Println("\033[33;1m\033[33;4m========== Etage 1 ==========\033[0m")
		fmt.Println()
		fmt.Println("\033[36;1mNavi \033[0m: Nous voici au premier étage.")
		scanner.Scan()
		scanner.Text()
		fmt.Println("\033[36;1mNavi \033[0m: Regarde ! Il y a un Bokoblin la-bas. Va-t'en debarasser. ")
		scanner.Scan()
		scanner.Text()
		tutorial_fight(p1, enemies[11])
		etage++
		money += 1
		fmt.Println("\033[36;1mNavi \033[0m: Heureusement que ce Bokoblin etait déja blessé.")
		scanner.Scan()
		scanner.Text()
		fmt.Println("\033[36;1mNavi \033[0m: Restons sur nos garde.")
		scanner.Scan()
		scanner.Text()
		return etage, money
	}
	return etage, money
}

func tutorial_fight(p1, enemy structure.Character) bool {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[33;1m\033[33;4m========== COMBAT ==========\033[0m")
	fmt.Println()
	enemy.Max_Hp = enemy.Hp
	for {
		enemy.Hp = 15

		fmt.Printf("\033[31;1m %v \033[0m\n", enemy.Name)
		fmt.Print("\033[97;1mHP :\033[0m ")
		for hp := 0; hp < enemy.Hp; hp++ {
			fmt.Print("I")
		}
		for hplost := enemy.Hp; hplost < enemy.Max_Hp; hplost++ {
			fmt.Print("_")
		}
		fmt.Printf(" %v / %v\n", enemy.Hp, enemy.Max_Hp)
		fmt.Println()

		// Insert Link HP
		fmt.Println()
		fmt.Println("\033[36;1mNavi \033[0m: Tu peux voir la barre de vie de l'ennemi en haut a gauche ainsi que son nom.")
		scanner.Scan()
		scanner.Text()
		fmt.Println("\033[36;1mNavi \033[0m: Lorsque la barre de vie de l'ennemi tombe a 0, tu remportes le combat.")
		scanner.Scan()
		scanner.Text()
		fmt.Println("\033[36;1mNavi \033[0m: L'ennemi est déja pas mal affaibli. Fini le pour de bon !")
		scanner.Scan()
		scanner.Text()
		fmt.Println()

		for enemy.Hp > 0 {
			fmt.Printf("\033[31;1m %v \033[0m\n", enemy.Name)
			fmt.Print("\033[97;1mHP :\033[0m ")
			for hp := 0; hp < enemy.Hp; hp++ {
				fmt.Print("I")
			}
			for hplost := enemy.Hp; hplost < enemy.Max_Hp; hplost++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", enemy.Hp, enemy.Max_Hp)
			fmt.Println()
			fmt.Println("\033[97;1mA. Attack		H. Heal\033[0m")

			fmt.Println()

			scanner.Scan()
			action := scanner.Text()
			switch action {
			case "A":
				enemy.Hp = enemy.Hp - p1.Str
				fmt.Printf("\033[97;3mVous attaquez \033[0m\033[97;1m%v\033[0m\033[97;3m et lui infliger \033[0m\033[97;1m%v\033[0m\033[97;3m degats !\033[0m\n", enemy.Name, p1.Str)
				fmt.Println()
			default:
				fmt.Println("\033[36;1mNavi \033[0m: Attaque l'ennemi en appuyant sur \033[97;1mA\033[0m !")
				fmt.Println()
			}
		}
		if enemy.Hp <= 0 {
			fmt.Println("\033[97;3mVous avez vaincu \033[0m", enemy.Name)
			scanner.Scan()
			scanner.Text()
			return true
		}
	}
}
