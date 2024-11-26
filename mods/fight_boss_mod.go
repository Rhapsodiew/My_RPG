package main

import (
	"mods/structure"
	"fmt"
)

func fight_boss(p1 *structure.Character, enemy structure.Character) bool {
	var text_pass rune
	fmt.Println("========== COMBAT ==========")
	fmt.Println()
	enemy.Max_Hp=enemy.Hp
	base_str := enemy.Str
	for {
		var end_turn bool
		end_turn = false
		var action int
		for !end_turn {
			// Enemy HP
			fmt.Printf("\033[33;1m %v \033[0m\n", enemy.Name)
			fmt.Print("HP : ")
			for hp := 0; hp < enemy.Hp; hp++ {
				fmt.Print("I")
			}
			for hplost := enemy.Hp; hplost < enemy.Max_Hp; hplost++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", enemy.Hp, enemy.Max_Hp)
			fmt.Println()
			
			// P1 HP

			fmt.Printf("\033[32;1m %v \033[0m\n", p1.Name)
			fmt.Print("HP : ")
			for hp := 0; hp < p1.Hp; hp++ {
				fmt.Print("I")
			}
			for hplost := p1.Hp; hplost < p1.Max_Hp; hplost++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", p1.Hp, p1.Max_Hp)
			fmt.Println()
			
			fmt.Println("1. Attack		2. Heal")
			fmt.Println("3. Magic 		4. Bag")
			fmt.Println()

			fmt.Scan(&action)
			switch action {
			case 1:
				enemy.Hp = enemy.Hp - p1.Str
				fmt.Printf("Vous attaquez %v et lui infliger %v degats !\n", enemy.Name, p1.Str)
				fmt.Println()
				end_turn = true
			case 2:
				p1.Hp += p1.Max_Hp / 2
				if p1.Hp > p1.Max_Hp {
					p1.Hp = p1.Max_Hp
				}
				fmt.Println("Vous vous soignez de ", int(p1.Max_Hp/2))
				fmt.Println()
				end_turn = true
			case 3:
				fmt.Println("\033[36;1mNavi \033[0m: Qu'est ce que tu essayes de faire ?")
			case 4:
				fmt.Println("\033[36;1mNavi \033[0m: Je te rappelle que l'on a égarer notre sac dans la forêt d'Hyrule ...")
			default:
				fmt.Println("\033[36;1mNavi \033[0m: Qu'est ce que tu essayes de faire ?")
				fmt.Println()
			}
		}
		
		if enemy.Hp <= enemy.Max_Hp/2 && enemy.Str == base_str{
			fmt.Printf("%s : Tu n'est donc pas arriver ici pour rien.\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Printf("%s : RWHAAAA\n",enemy.Name)
			fmt.Println("\033[36;1mNavi \033[0m: Oh non, son corp se regenere et devient encore plus grand qu'avant")
			fmt.Println()
			enemy.Hp += enemy.Max_Hp /4
			enemy.Str += 5
			
		}

		if enemy.Hp <= 0 {
			fmt.Printf("%s : Keugh\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Printf("%s : ...\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Printf("%s : Ne t'inquiete pas ...\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Printf("%s : Je reviendrais plus fort que jamais ...\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Printf("%s : Et Hyrule sombrera dans la terreur !\n",enemy.Name)
			fmt.Scan(&text_pass)
			fmt.Println("Vous avez vaincu ", enemy.Name)
			fmt.Scan(&text_pass)
			return true
		}
		p1.Hp = p1.Hp - enemy.Str
		fmt.Println(enemy.Name, " vous attaque et inflige ", enemy.Str, " degat !")
		fmt.Println()

		if p1.Hp <= 0 {
			fmt.Println("Vous etes mort")
			fmt.Scan(&text_pass)
			fmt.Printf("%s : Idiot, tu pensais être de taille face a MOI, GANON, AHAHAHAH !!!!\n",enemy.Name)
			fmt.Println()
			return false
		}
	}
}
