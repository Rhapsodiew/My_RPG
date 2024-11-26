package main

import (
	"bufio"
	"fmt"
	characteristics "mods/basic_characteristics"
	gamecustom "mods/basic_game_customization"
	bettercombatoption "mods/better_combat_option"
	"mods/structure"
	"os"
)

func main() {
	var mods []string
	scanner := bufio.NewScanner(os.Stdin)
	//TITLE SCREEN + SETTINGS
	var lunch_game bool
	var difficulty int
	var max_etage int
	for !lunch_game {
		Ascii_art_castle()
		gamecustom.Title_screen()
		scanner.Scan()
		start := scanner.Text()
		mods = gamecustom.Check_mods(start, mods)
		start, difficulty, max_etage = gamecustom.Check_start(start)
		if start == "Q" {
			return
		} else if start == "N" {
			lunch_game = true
		}
	}
	fmt.Println("Mods : ", mods)
	var bag []structure.Item
	enemies, bosses := change_mob_stats(difficulty)
	xp := 0
	money := 12
	boss := chose_random_boss(bosses)
	p1 := chose_random_player()
	etage := 1

	// ================== START ======================
	scanner.Scan()
	scanner.Text()
	Introduction(p1, boss)

	// ========= TUTORIEL ============

	fmt.Println("\033[33;1m-> \033[97;1mVoulez vous faire le tutoriel ?\033[0m   y/n")
	scanner.Scan()
	tuto := scanner.Text()
	etage, money = tutoriel(tuto, p1, enemies, etage, money)

	//================ ETAGE TOWER =================

	for etage < max_etage+1 {
		stage_enemies := chose_random_enemies(etage, enemies)
		// fmt.Println(stage_enemies)

		result := fight_mod(&p1, stage_enemies, etage, max_etage, bag, xp, mods)
		if !result {
			fmt.Println("Vous avez echouer !")
			return
		} else {
			mod_activate := gamecustom.Is_mod_activate(mods, "Level_And_Experience")
			if mod_activate {
				xp = characteristics.Get_xp(xp)
				xp = characteristics.Levelup(&p1, xp)
			}

			mod_activate = gamecustom.Is_mod_activate(mods, "Inventory")
			if mod_activate {
				bag = bettercombatoption.Get_item(bag)
			}

			fmt.Println("\033[36;1mNavi \033[0m: Bien jouer, passons au prochaine etage !")
			scanner.Scan()
			scanner.Text()
			etage++
			mod_activate = gamecustom.Is_mod_activate(mods, "Random_Game_Events")
			if mod_activate {
				secret := gamecustom.Is_it_a_secret_room(etage)
				trap := gamecustom.Which_room()
				gamecustom.Trap_room(secret, trap, &p1, &money)
				if secret{
					Traps(trap)
				}
			}
			money += 1
		}

		// ==================== ETAGE BOSS ========================

		if etage%10 == 0 {
			// Boss_intro(boss)
			// result := fight_boss(&p1, boss)
			result := fight_mod(&p1, boss, etage, max_etage, bag, xp, mods)
			if !result {
				fmt.Println("\033[97;3mVous avez echouer !\033[0m")
				return
			} else {
				if etage == max_etage {
					fmt.Println("\033[36;1mNavi \033[0m: Tu as sauver Hyrule de ce monstre !")
					fmt.Println("\033[36;1mNavi \033[0m: Hyrule pourra enfin connaitre la paix.")
					etage += 1
				} else {
					mod_activate := gamecustom.Is_mod_activate(mods, "Level_And_Experience")
					if mod_activate {
						xp = characteristics.Get_xp(xp)
						xp = characteristics.Levelup(&p1, xp)
					}
					fmt.Println("\033[36;1mNavi \033[0m: Il te reste encore quelque étage a franchir pour arrive au sommet")
					fmt.Println("\033[36;1mNavi \033[0m: Continue comme ca !")
					scanner.Scan()
					scanner.Text()
					boss = chose_random_boss(bosses)
					etage += 1
				}

			}
		}
	}
}
