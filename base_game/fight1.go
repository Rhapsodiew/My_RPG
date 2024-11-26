package main

import "fmt"

func fight1(p1, enemy Character) bool{

	fmt.Println("========== COMBAT ==========")
	fmt.Println()
	for {
		var text_pass rune
		enemy.Health = 15
		fmt.Printf("%v \n", enemy.Name)
		fmt.Print("HP : ")
		for hp := 0; hp < enemy.Health; hp++ {
			fmt.Print("I")
		}
		for hplost := enemy.Health; hplost < enemy.Max_health; hplost++ {
			fmt.Print("_")
		}
		fmt.Printf(" %v / %v\n", enemy.Health, enemy.Max_health)
		fmt.Println()

		// Insert Link HP
		fmt.Println()
		var pass_text rune
		fmt.Println("Navi : Tu peux voir la barre de vie de l'ennemi en haut a gauche ainsi que son nom.")
		fmt.Scan(&pass_text)
		fmt.Println("Navi : Lorsque la barre de vie de l'ennemi tombe a 0, tu remportes le combat.")
		fmt.Scan(&pass_text)
		fmt.Println("Navi : L'ennemi est déja pas mal affaibli. Fini le pour de bon !")
		fmt.Scan(pass_text)
		fmt.Println()

		for enemy.Health > 0 {
			fmt.Printf("%v \n", enemy.Name)
			fmt.Print("HP : ")
			for hp := 0; hp < enemy.Health; hp++ {
				fmt.Print("I")
			}
			for hplost := enemy.Health; hplost < enemy.Max_health; hplost++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", enemy.Health, enemy.Max_health)
			fmt.Println()
			fmt.Println("1. Attack		2. Heal")
			fmt.Println("3. Magic 		4. Bag")

			fmt.Println()

			
			var action int
			fmt.Scan(&action)
			switch action {
			case 1:
				enemy.Health = enemy.Health - p1.Strength
				fmt.Printf("Vous attaquez %v et lui infliger %v degats !\n", enemy.Name, p1.Strength)
				fmt.Println()
			default:
				fmt.Println("Navi : Attaque l'ennemi en appuyant sur 1 !")
				fmt.Println()
			}
		}
		
		if enemy.Health <= 0 {
			fmt.Println("Vous avez vaincu ", enemy.Name)
			fmt.Scan(&text_pass)
			return true
		}
		p1.Health = p1.Health - enemy.Strength
		fmt.Println(enemy.Name, " vous attaque et inflige ", enemy.Strength, " degat !")

		if p1.Health <= 0 {
			fmt.Println("Vous etes mort")
			fmt.Scan(&text_pass)
			return false
		}
	}
}
