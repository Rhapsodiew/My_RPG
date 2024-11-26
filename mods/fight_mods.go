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

func fight_mod(p1 *structure.Character, enemy structure.Character, etage int, max_etage int, bag []structure.Item, xp int, mods []string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	var isboss bool
	var phase2 bool
	fmt.Println()
	fmt.Printf("\033[33;1m\033[33;4m========== Etage %d ==========\033[0m\n", etage)
	scanner.Scan()
	scanner.Text()
	fmt.Println()
	if etage%10 == 0 {
		Door()
		scanner.Scan()
		scanner.Text()
		Boss_intro(enemy)
		isboss = true
	}else if etage == 3{
		Etage_3()
	}

	fmt.Println("\033[33;1m\033[33;4m========== COMBAT ==========\033[0m")
	fmt.Println()
	enemy.Max_Hp = enemy.Hp
	var used_special bool

	for {
		var first string
		var action string
		// WHO IS FASTER
		mod_activate_more_statistics := gamecustom.Is_mod_activate(mods, "More_Statistics")
		if mod_activate_more_statistics {
			isfaster := bettercombatoption.Is_faster(p1, &enemy)
			if isfaster {
				first = p1.Name
			} else {
				first = enemy.Name
			}
		} else {
			first = p1.Name
		}
		if isboss {
			if enemy.Hp <= enemy.Max_Hp/2 && !phase2 {
				Boss_phase2(enemy)
				phase2 = true
			}
		}

		// Enemy HP
		fmt.Printf("\033[31;1m %v \033[0m\n", enemy.Name)
		fmt.Print("\033[97;1m HP :\033[0m  ")
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
		fmt.Print("\033[97;1m HP :\033[0m  ")
		for hp := 0; hp < p1.Hp; hp++ {
			fmt.Print("I")
		}
		for hplost := p1.Hp; hplost < p1.Max_Hp; hplost++ {
			fmt.Print("_")
		}
		fmt.Printf(" %v / %v\n", p1.Hp, p1.Max_Hp)

		// P1 MP
		mod_activate_magic_skills := gamecustom.Is_mod_activate(mods, "Magic_Skills")
		if mod_activate_magic_skills {
			fmt.Print("\033[36;1m Mana \033[0m : ")
			for basemana := 0; basemana < p1.Mp; basemana++ {
				fmt.Print("O")
			}
			for lostmana := p1.Mp; lostmana < p1.Max_Mp; lostmana++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", p1.Mp, p1.Max_Mp)
		}
		// P1 XP
		mod_activate_level_and_experience := gamecustom.Is_mod_activate(mods, "Level_And_Experience")
		if mod_activate_level_and_experience {
			fmt.Print("\033[33;1m Xp \033[0m : ")
			for basexp := 0; basexp < xp; basexp++ {
				fmt.Print("|")
			}
			for missingxp := xp; missingxp < 100; missingxp++ {
				fmt.Print("_")
			}
			fmt.Printf(" %v / %v\n", xp, 100)
		}
		fmt.Println()

		//OPTION
		fmt.Print("\033[97;1mA. Action\033[0m")
		if mod_activate_magic_skills {
			bettercombatoption.Skills_option() // MAGIC SKILLS OPTION
		} else {
			fmt.Print("		\033[97;1mH. Heal\033[0m\n") // HEAL OPTION
		}

		mod_activate_better_combat_option := gamecustom.Is_mod_activate(mods, "Better_Combat_Option") // DEFEND / ESCAPE OPTION
		if mod_activate_better_combat_option {
			bettercombatoption.Better_combat_option()
		}
		mod_activate_basic_characteristics_2 := gamecustom.Is_mod_activate(mods, "Basic_Characteristics_2") // SPECIAL MOVE OPTION
		if mod_activate_basic_characteristics_2 {
			characteristics.Option_special_move()
		}
		mod_activate_inventory := gamecustom.Is_mod_activate(mods, "Inventory") // INVENTORY OPTION
		if mod_activate_inventory {
			bettercombatoption.Option_Inventory()
		}
		mod_activate_basic_characteristics := gamecustom.Is_mod_activate(mods, "Basic_Characteristics")
		if mod_activate_basic_characteristics {
			characteristics.Option_character()
		}

		scanner.Scan()
		action = scanner.Text()
		var dmg1 float64
		var dmg2 float64
		var new_dmg1 float64
		var new_dmg2 float64
		if action == "A" || action == "Attack" {
			if first == p1.Name {
				if mod_activate_more_statistics { // CALCULATE DMG P1
					dmg1 = bettercombatoption.Calculate_type_attk(p1)
					dmg1 = bettercombatoption.Crit(p1, dmg1)
					dmg1 = bettercombatoption.Calculate_def_res(p1, &enemy, dmg1)
				} else {
					dmg1 = float64(p1.Str)
				}

				if mod_activate_basic_characteristics {
					new_dmg1 = characteristics.Calculate_race_class(p1, &enemy, dmg1)
					characteristics.Is_Effective(dmg1, new_dmg1, p1)
				} else {
					new_dmg1 = dmg1
				}
				enemy.Hp -= int(new_dmg1)
				fmt.Println("\033[97;3mVous infligez \033[97;1m", int(new_dmg1), "\033[0m\033[97;3m degats\033[0m")
				if enemy.Hp <= 0 {
					fmt.Println()
					if isboss {
						if etage == max_etage {
							Boss_end(enemy)
						} else {
							Boss_mid_end(enemy)
						}
					} else {
						fmt.Printf("\033[97;3mVous avez vaincu \033[97;1m%s\033[0m \n", enemy.Name)
						scanner.Scan()
						scanner.Text()
					}

					return true
				}

				if mod_activate_more_statistics { // CALCULATE DMG ENEMY
					dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
					dmg2 = bettercombatoption.Crit(&enemy, dmg2)
					dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
				} else {
					dmg2 = float64(enemy.Str)
				}

				if mod_activate_basic_characteristics {
					new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
					characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
				} else {
					new_dmg2 = dmg2
				}

				if mod_activate_more_statistics {
					dodge := bettercombatoption.Dodge(p1, &enemy)
					if dodge {
						new_dmg2 = 0
					}
				}

				p1.Hp -= int(new_dmg2)
				fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
				fmt.Println()
				scanner.Scan()
				scanner.Text()
				if p1.Hp <= 0 {
					fmt.Println("\033[97;3mVous etes mort\033[0m")
					scanner.Scan()
					scanner.Text()
					if isboss {
						Boss_u_die(enemy)
					}
					return false
				}

			} else {
				dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
				dmg2 = bettercombatoption.Crit(&enemy, dmg2)
				dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
				if mod_activate_basic_characteristics {
					new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
					characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
				} else {
					new_dmg2 = dmg2
				}

				p1.Hp -= int(new_dmg2)
				fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
				if p1.Hp <= 0 {
					fmt.Println("\033[97;3mVous etes mort\033[0m")
					scanner.Scan()
					scanner.Text()
					if isboss {
						Boss_u_die(enemy)
					}
					return false
				}
				dmg1 := bettercombatoption.Calculate_type_attk(p1)
				dmg1 = bettercombatoption.Crit(p1, dmg1)
				dmg1 = bettercombatoption.Calculate_def_res(p1, &enemy, dmg1)

				if mod_activate_basic_characteristics {
					new_dmg1 := characteristics.Calculate_race_class(p1, &enemy, dmg1)
					characteristics.Is_Effective(dmg1, new_dmg1, p1)
				} else {
					new_dmg1 = dmg1
				}

				dodge := bettercombatoption.Dodge(&enemy, p1)
				if dodge {
					new_dmg1 = 0
				}
				enemy.Hp -= int(new_dmg1)
				fmt.Printf("\033[97;3mVous infligez \033[97;1m%v\033[0m\033[97;3m degats\n", int(new_dmg1))
				scanner.Scan()
				scanner.Text()
				if enemy.Hp <= 0 {
					if isboss {
						if etage == max_etage {
							Boss_end(enemy)
						} else {
							Boss_mid_end(enemy)
						}
					} else {
						fmt.Printf("\033[97;3mVous avez vaincu \033[97;1m%s\033[0m\n", enemy.Name)
						scanner.Scan()
						scanner.Text()
					}
					return true
				}
			}

		} else if action == "S" || action == "Skills" {
			if mod_activate_magic_skills {
				var dead_after_spell bool
				var valide bool
				p1, enemy, dead_after_spell, valide = bettercombatoption.Magic_skills_list(p1, enemy)
				if dead_after_spell {
					if isboss {
						if etage == max_etage {
							Boss_end(enemy)
						} else {
							Boss_mid_end(enemy)
						}
					} else {
						fmt.Printf("\033[97;3mVous avez vaincu \033[97;1m%s\033[0m\n", enemy.Name)
						scanner.Scan()
						scanner.Text()
					}

					return true
				}
				if valide {
					if mod_activate_more_statistics {
						dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
						dmg2 = bettercombatoption.Crit(&enemy, dmg2)
						dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
					} else {
						dmg2 = float64(enemy.Str)
					}

					if mod_activate_basic_characteristics {
						new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
						characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
					} else {
						new_dmg2 = dmg2
					}

					p1.Hp -= int(new_dmg2)
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
					scanner.Scan()
					scanner.Text()
					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort\033[")
						scanner.Scan()
						scanner.Text()
						if isboss {
							Boss_u_die(enemy)
						}
						return false
					}
				}
			}

		} else if action == "D" || action == "Defend" {
			if mod_activate_better_combat_option {
				if mod_activate_more_statistics {
					dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
					dmg2 = bettercombatoption.Crit(&enemy, dmg2)
					dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
				} else {
					dmg2 = float64(enemy.Str)
				}

				if mod_activate_basic_characteristics {
					new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
					characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
				} else {
					new_dmg2 = dmg2
				}

				if first == p1.Name {
					fmt.Println("\033[97;3mVous vous mettez en garde !\033[0m")
					fmt.Println()
					defend_dmg := bettercombatoption.Defend(new_dmg2)
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(defend_dmg))
					p1.Hp -= int(defend_dmg)
					scanner.Scan()
					scanner.Text()
					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort\033[0m")
						scanner.Scan()
						scanner.Text()
						if isboss {
							Boss_u_die(enemy)
						}
						return false
					}
				} else {
					fmt.Println("\033[97;3mVous n'etes pas assez rapide pour vous proteger des l'ennemi\033[0m")
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
					p1.Hp -= int(new_dmg2)
					scanner.Scan()
					scanner.Text()
					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort\033[0m")
						scanner.Scan()
						scanner.Text()
						if isboss {
							Boss_u_die(enemy)
						}
						return false
					}
				}
			}

		} else if action == "E" || action == "Escape" {
			if !isboss {
				if mod_activate_better_combat_option {
					if mod_activate_more_statistics {
						dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
						dmg2 = bettercombatoption.Crit(&enemy, dmg2)
						dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
					} else {
						dmg2 = float64(enemy.Str)
					}

					if mod_activate_basic_characteristics {
						new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
						characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
					} else {
						new_dmg2 = dmg2
					}

					p1.Hp = p1.Hp - int(new_dmg2)*5
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))

					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort en essayant de fuir !\033[0m")
						scanner.Scan()
						scanner.Text()
						return false
					} else {
						fmt.Println("\033[97;3mVous avez réussi a vous enfuir !\033[0m")
						scanner.Scan()
						scanner.Text()
						return true
					}
				}
			}
		} else if action == "C" || action == "Character" {
			if mod_activate_basic_characteristics {
				characteristics.Show_stats(p1)
				if mod_activate_basic_characteristics_2 {
					characteristics.Show_stats_2(p1)
				}
				scanner.Scan()
				scanner.Text()
			}

		} else if action == "B" || action == "Bag" {
			if mod_activate_inventory {
				var valide bool
				p1, bag, valide = bettercombatoption.Show_inventory(p1, bag)
				if valide {
					if mod_activate_more_statistics {
						dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
						dmg2 = bettercombatoption.Crit(&enemy, dmg2)
						dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
					} else {
						dmg2 = float64(enemy.Str)
					}

					if mod_activate_basic_characteristics {
						new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
						characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
					} else {
						new_dmg2 = dmg2
					}

					p1.Hp -= int(new_dmg2)
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
					scanner.Scan()
					scanner.Text()
					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort\033[0m")
						scanner.Scan()
						scanner.Text()
						if isboss {
							Boss_u_die(enemy)
						}
						return false
					}
				}
			}

		} else if action == "Sp" || action == "Special" {
			if mod_activate_basic_characteristics_2 {
				used_special = characteristics.Special_move(used_special, &enemy)
			}

		} else if action == "H" || action == "Heal" {
			if !mod_activate_magic_skills {
				if p1.Hp != p1.Max_Hp {
					hp := p1.Hp
					p1.Hp += p1.Max_Hp / 2
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					fmt.Printf("\033[97;3mVous vous etes soigné de \033[97;1m%v\033[0m", p1.Hp-hp)
					fmt.Println()
					if mod_activate_more_statistics {
						dmg2 = bettercombatoption.Calculate_type_attk(&enemy)
						dmg2 = bettercombatoption.Crit(&enemy, dmg2)
						dmg2 = bettercombatoption.Calculate_def_res(&enemy, p1, dmg2)
					} else {
						dmg2 = float64(enemy.Str)
					}

					if mod_activate_basic_characteristics {
						new_dmg2 = characteristics.Calculate_race_class(&enemy, p1, dmg2)
						characteristics.Is_Effective(dmg2, new_dmg2, &enemy)
					} else {
						new_dmg2 = dmg2
					}

					p1.Hp -= int(new_dmg2)
					fmt.Printf("\033[97;3m\033[97;1m%s\033[0m\033[97;3m vous inflige \033[97;1m%v\033[0m\033[97;3m degats\033[0m \n", enemy.Name, int(new_dmg2))
					scanner.Scan()
					scanner.Text()
					if p1.Hp <= 0 {
						fmt.Println("\033[97;3mVous etes mort\033[0m")
						scanner.Scan()
						scanner.Text()
						if isboss {
							Boss_u_die(enemy)
						}
						return false
					}
				} else {
					fmt.Println("\033[97;3mVous avez toute votre vie\033[0m")
					scanner.Scan()
					scanner.Text()
				}

			}
		}
	}
}

//CHANGE DMG1 and DMG2 TO NEW_DMG1 NEW_DMG2
// 	if fuite {
// 		if first == p1.Name {
// 			new_dmg2 = new_dmg2 * 5
// 			p1.Hp = p1.Hp - int(new_dmg2)
// 			if p1.Hp <= 0 {
// 				fmt.Println("Vous etes mort en essayant de fuir !")
// 				fmt.Scan(&text_pass)
// 				return false
// 			} else {
// 				return true
// 			}
// 		} else {
// 			new_dmg1 = new_dmg1 * 5
// 			p1.Hp = p1.Hp - int(new_dmg1)
// 			if p1.Hp <= 0 {
// 				fmt.Println("Vous etes mort en essayant de fuir !")
// 				fmt.Scan(&text_pass)
// 				return false
// 			} else {
// 				return true
// 			}
// 		}

// 	}

// 	if first == p1.Name {
// 		if attack {
// 			enemy.Hp = enemy.Hp - int(new_dmg1)
// 			fmt.Println("Vous attaquez et infligez ", int(new_dmg1), " degat !")
// 			fmt.Println()
// 			if enemy.Hp <= 0 {
// 				fmt.Println("Vous avez vaincu ", enemy.Name)
// 				fmt.Scan(&text_pass)
// 				return true
// 			}
// 		} else if heal {
// 			p1.Hp += p1.Max_Hp / 2
// 			if p1.Hp > p1.Max_Hp {
// 				p1.Hp = p1.Max_Hp
// 			}
// 			fmt.Println("Vous vous soignez de ", int(p1.Max_Hp/2))
// 			fmt.Println()
// 		} else if defend {
// 			new_dmg2 = new_dmg2 / 2
// 		}
// 		p1.Hp = p1.Hp - int(new_dmg2)
// 		fmt.Println(enemy.Name, " vous attaque et inflige ", int(new_dmg2), " degat !")
// 		fmt.Println()
// 		if p1.Hp <= 0 {
// 			fmt.Println("Vous etes mort")
// 			fmt.Scan(&text_pass)
// 			return false
// 		}
// 	} else {
// 		p1.Hp = p1.Hp - int(new_dmg1)
// 		fmt.Println(enemy.Name, " vous attaque et inflige ", int(new_dmg1), " degat !")
// 		fmt.Println()
// 		if p1.Hp <= 0 {
// 			fmt.Println("Vous etes mort")
// 			fmt.Scan(&text_pass)
// 			return false
// 		}
// 		if attack {
// 			enemy.Hp = enemy.Hp - int(new_dmg2)
// 			fmt.Println("Vous attaquez et infligez ", int(new_dmg2), " degat !")
// 			fmt.Println()
// 			if enemy.Hp <= 0 {
// 				fmt.Println("Vous avez vaincu ", enemy.Name)
// 				fmt.Scan(&text_pass)
// 				return true
// 			}
// 		} else if heal {
// 			p1.Hp += p1.Max_Hp / 2
// 			if p1.Hp > p1.Max_Hp {
// 				p1.Hp = p1.Max_Hp
// 			}
// 			fmt.Println("Vous vous soignez de ", int(p1.Max_Hp/2))
// 			fmt.Println()
// 		} else if defend {
// 			fmt.Println("Vous n'etes pas assez rapide pour vous proteger des l'ennemi")
// 		}
// 	// 	}
// 	}
// }
