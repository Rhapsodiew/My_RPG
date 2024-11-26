package bettercombatoption

import (
	"bufio"
	"fmt"
	"mods/structure"
	"os"
	"strings"
)

func Skills_option() {
	fmt.Print("		\033[97;1mS. Skills\033[0m\n")
}

func Magic_skills_list(p1 *structure.Character, enemy structure.Character) (*structure.Character, structure.Character, bool, bool) {
	scanner := bufio.NewScanner(os.Stdin)

	spells := structure.Read_spells()
	classes := structure.Read_classes()
	fmt.Println("\033[31;1m\033[31;4mSPELLS\033[0m")			// SHOW SPELL AND ADD THEM TO SLICE
	fmt.Println("---------")
	fmt.Printf("\033[32;1m%s\033[0m", spells[3].Name)
	fmt.Printf("\033[32;1m	%s\033[0m", spells[4].Name)
	fmt.Printf("\033[31;1m	%s\033[0m\n", spells[5].Name)

	usable_spell := []string{spells[3].Name, spells[4].Name, spells[5].Name}
	for _, class := range classes {
		if p1.Class == class.Id {
			for _, spell := range spells {
				if strings.Contains(spell.Class, strings.ToLower(class.Name)) {
					fmt.Printf("\033[97;1m%s\033[0m		", spell.Name)
					usable_spell = append(usable_spell, spell.Name)
				}
			}
		}
	}
	fmt.Println()
	fmt.Println()

	scanner.Scan()
	answer := scanner.Text()
	valide := false
	// Spells_use(answer,usable_spell,p1)
	for i := 0; i < len(usable_spell); i++ {		// USE OF EVERY SPELL
		if answer == usable_spell[i] {
			if answer == "Cheat Heal" {
				if p1.Hp != p1.Max_Hp {
					if p1.Mp >= spells[3].Cost {
						fmt.Println("\033[97;1m\033[97;3mCHEAT HEAL !!\033[0m")
						p1.Mp -= spells[3].Cost
						hp := p1.Hp
						p1.Hp += 100
						if p1.Hp > p1.Max_Hp {
							p1.Hp = p1.Max_Hp
						}
						fmt.Printf("\033[97;3mVous vous etes soigné de \033[97;1m%v\033[0m\n", p1.Hp-hp)

						// fmt.Println(p1.Hp,p1.Mp)
						valide = true
					} else {
						fmt.Println("\033[97;4mPas assez de mana\033[0m")
						scanner.Scan()
						scanner.Text()
					}
				} else {
					fmt.Println("\033[97;4mVous avec deja toute votre vie\033[0m")
					scanner.Scan()
					scanner.Text()
				}
			} else if answer == "Cheat Restore" {
				fmt.Println("\033[97;1m\033[97;3mCHEAT RESTORE !!\033[0m")

				mp := p1.Mp
				p1.Mp += 100
				if p1.Mp > p1.Max_Mp {
					p1.Mp = p1.Max_Mp
				}
				fmt.Printf("\033[97;3mVous avez recupé \033[97;1m%v\033[0m\033[97;3m point de mana\n", p1.Mp-mp)

				if p1.Mp > p1.Max_Mp {
					p1.Mp = p1.Max_Mp
				}
				valide = true
			} else if answer == "Cheat Fireball" {
				if p1.Mp >= spells[5].Cost {
					fmt.Println("\033[97;1m\033[97;3mCHEAT FIREBALL !!\033[0m")

					p1.Mp -= spells[5].Cost
					enemy.Hp -= spells[5].Dmg
					fmt.Printf("\033[97;3mVous infligez \033[97;1m%v\033[0m\033[97;3m degats\n", int(spells[5].Dmg))
					scanner.Scan()
					scanner.Text()
					if enemy.Hp <= 0 {
						return p1, enemy, true, valide
					}
					valide = true
				} else {
					fmt.Println("\033[97;4mPas assez de mana\033[0m")
					scanner.Scan()
					scanner.Text()
				}
			} else if answer == "Heal" {
				if p1.Hp != p1.Max_Hp {
					if p1.Mp >= spells[1].Cost {
						fmt.Println("\033[97;1m\033[97;3mHEAL !!\033[0m")

						hp := p1.Hp
						p1.Mp -= spells[1].Cost
						p1.Hp += 20
						if p1.Hp > p1.Max_Hp {
							p1.Hp = p1.Max_Hp
						}
						fmt.Printf("\033[97;3mVous vous etes soigné de \033[97;1m%v\033[0m\n", p1.Hp-hp)

						// fmt.Println(p1.Hp,p1.Mp)
						valide = true
					} else {
						fmt.Println("\033[97;4mPas assez de mana\033[0m")
						scanner.Scan()
						scanner.Text()
					}
				} else {
					fmt.Println("\033[97;4mVous avez deja toute votre vie\033[0m")
					scanner.Scan()
					scanner.Text()
				}
			} else if answer == "Heal II" {
				if p1.Hp != p1.Max_Hp {
					if p1.Mp >= spells[2].Cost {
						fmt.Println("\033[97;1m\033[97;3mHEAL !!\033[0m")

						hp := p1.Hp
						p1.Mp -= spells[2].Cost
						p1.Hp += 50
						if p1.Hp > p1.Max_Hp {
							p1.Hp = p1.Max_Hp
						}
						fmt.Printf("\033[97;3mVous vous etes soigné de \033[97;1m%v\033[0m\n", p1.Hp-hp)

						// fmt.Println(p1.Hp,p1.Mp)
						valide = true
					} else {
						fmt.Println("\033[97;4mPas assez de mana\033[0m")
						scanner.Scan()
						scanner.Text()
					}
				} else {
					fmt.Println("\033[97;4mVous avez deja toute votre vie\033[0m")
					scanner.Scan()
					scanner.Text()
				}
			}
		}
	}
	return p1, enemy, false, valide
}

// func Spells_use(answer string, usable_spell []string, p1 *structure.Character) {
// 	spells := structure.Read_spells()

// 	for _, spell_list := range usable_spell{
// 		if answer == spell_list{
// 			if answer == "Cheat Heal"{

// 			}
// 		}
// 	}
// }
