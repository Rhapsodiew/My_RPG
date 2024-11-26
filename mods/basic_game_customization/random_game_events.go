package gamecustom

import (
	"bufio"
	"fmt"
	"math/rand"
	"mods/structure"
	"os"
	"time"
)

func Is_it_a_secret_room(etage int) bool {
	var secret bool
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	random := (rand.Intn(max+1) + min)
	if random <= 35 {
		secret = true
	}
	if etage%10 == 0 {
		secret = true
	}
	return secret
}

func Which_room() structure.Traps {
	var chose_trap []structure.Traps
	var chosen_trap structure.Traps
	traps := structure.Read_traps()

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	random_trap := (rand.Intn(max) + min)
	if random_trap <= 30 {
		for i := 0; i < len(traps); i++ {
			if traps[i].Rarity == 1 {
				chose_trap = append(chose_trap, traps[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		random_trap = (rand.Intn(len(chose_trap)) + min)
		chosen_trap = chose_trap[random_trap-1]
	} else if random_trap > 30 && random_trap <= 60 {
		for i := 0; i < len(traps); i++ {
			if traps[i].Rarity == 2 {
				chose_trap = append(chose_trap, traps[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		random_trap = (rand.Intn(len(chose_trap)) + min)
		chosen_trap = chose_trap[random_trap-1]
	} else {
		chosen_trap = structure.Traps{Name: "Treasure Room", Requirement: "Gain_gold", Rarity: 0}
	}
	return chosen_trap
}

func Trap_room(secret bool, trap structure.Traps, p1 *structure.Character,money *int) {
	scanner := bufio.NewScanner(os.Stdin)
	var action string
	if secret {
		// SPIKE TRAP
		if trap.Name == "Pike Wall" {
			for action != "L" && action != "Leave" {
				fmt.Println("\033[97;3mVous entrez dans une salle remplis de piege !\033[0m")
				fmt.Println("\033[97;3mPour en sortir indemne vous devez posseder plus de \033[0m\033[97;1m10\033[0m\033[97;3m de Force\033[0m")
				fmt.Println()
				fmt.Println("\033[97;1mL. Leave\033[0m")
				scanner.Scan()
				action = scanner.Text()
				fmt.Println()
				if action == "L" || action == "Leave" {
					if p1.Str < 10 {
						rand.Seed(time.Now().UnixNano())
						min := p1.Max_Hp / 20
						max := p1.Max_Hp * 3 / 20
						lost_hp := (rand.Intn(max) + min)
						p1.Hp -= lost_hp
						fmt.Printf("\033[97;3mVous avez perdu \033[0m\033[97;1m%v\033[0m\033[97;3m HP\033[0m\n", lost_hp)
					} else {
						fmt.Println("\033[97;3mVous passer sans probleme\033[0m")
						*money += 1
					}
				}
			}
		// CODE TRAP
		} else if trap.Name == "Secret Code" {
			for action != "L" && action != "Leave" {
				fmt.Println("\033[97;3mVous entrez dans une salle mystérieuse !\033[0m")
				fmt.Println("\033[97;3mPour en sortir indemne vous devez posseder plus de \033[0m\033[97;1m10\033[0m \033[97;3mde Inteligence\033[0m")
				fmt.Println()
				fmt.Println("\033[97;1mL. Leave\033[0m")
				scanner.Scan()
				action = scanner.Text()
				if action == "L" || action == "Leave" {
					if p1.Int < 10 {
						rand.Seed(time.Now().UnixNano())
						min := p1.Max_Hp / 20
						max := p1.Max_Hp * 3 / 20
						lost_hp := (rand.Intn(max+1-min) + min)
						p1.Hp -= lost_hp
						fmt.Printf("\033[97;3mVous avez perdu \033[0m\033[97;1m%v\033[0m\033[97;3m HP\033[0m\n", lost_hp)
					} else {
						fmt.Println("\033[97;3mVous passer sans probleme\033[0m")
						*money +=1
					}
				}
			}
		// TREASURE ROOM
		} else if trap.Name == "Treasure Room" {
			for action != "L" && action != "Leave" {
				rand.Seed(time.Now().UnixNano())
				min := 3
				max := 5
				gold := (rand.Intn(max+1-min) + min)
				fmt.Println("\033[97;3mVous entrez dans une salle au trésor !\033[0m")
				fmt.Printf("\033[97;3mVous y trouvez \033[0m\033[97;1m%v\033[0m\033[97;3m piece\n",gold)
				fmt.Println()
				fmt.Println("\033[97;1mL. Leave\033[0m")
				scanner.Scan()
				action = scanner.Text()
			}
		}
	}
}
