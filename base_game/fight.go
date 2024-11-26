package main

import "fmt"

func fight(p1 *Character, enemy Character) bool {
	var text_pass rune
	fmt.Println("========== COMBAT ==========")
	fmt.Println()

	for {
		var end_turn bool
		end_turn = false
		var action int
		for end_turn == false {
			// Enemy HP

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
			
			// Link HP

			fmt.Printf("%v \n", p1.Name)
			fmt.Print("HP : ")
			for hp := 0; hp < p1.Health; hp++ {
				fmt.Print("I")
			}
			for hplost := p1.Health; hplost < p1.Max_health; hplost++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", p1.Health, p1.Max_health)
			fmt.Println()
			
			fmt.Println("1. Attack		2. Heal")
			fmt.Println("3. Magic 		4. Bag")
			fmt.Println()

			fmt.Scan(&action)
			switch action {
			case 1:
				enemy.Health = enemy.Health - p1.Strength
				fmt.Printf("Vous attaquez %v et lui infliger %v degats !\n", enemy.Name, p1.Strength)
				fmt.Println()
				end_turn = true
			case 2:
				p1.Health += p1.Max_health / 2
				if p1.Health > p1.Max_health {
					p1.Health = p1.Max_health
				}
				fmt.Println("Vous vous soignez de ", int(p1.Max_health/2))
				fmt.Println()
				end_turn = true
			case 3:
				fmt.Println("Navi : Qu'est ce que tu essayes de faire ?")
			case 4:
				fmt.Println("Navi : Je te rappelle que l'on a égarer notre sac dans la forêt d'Hyrule ...")
			default:
				fmt.Println("Navi : Qu'est ce que tu essayes de faire ?")
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
		fmt.Println()

		if p1.Health <= 0 {
			fmt.Println("Vous etes mort")
			fmt.Scan(&text_pass)
			return false
		}
	}
}
