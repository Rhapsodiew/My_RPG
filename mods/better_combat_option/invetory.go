package bettercombatoption

import (
	"bufio"
	"fmt"
	"math/rand"
	"mods/structure"
	"os"
	"time"
)

func Option_Inventory() {
	fmt.Println("\033[97;1mB. Bag\033[0m")

}

func Show_inventory(p1 *structure.Character, bag []structure.Item) (*structure.Character,[]structure.Item,bool){
	scanner := bufio.NewScanner(os.Stdin)
	var valide bool
	fmt.Println("\033[32;1m\033[32;4mInventaire :\033[0m")
	for _, item := range bag {
		fmt.Println(item.Name)
	}
	fmt.Println("\033[32;4mQuel objet voulez vous utiliser ?\033[0m")		// SHOW INVENTORY AND USABLE ITEM
	scanner.Scan()															// IF USE ITEM -> PASS TOUR + ENEMY ATTACK
	answer := scanner.Text()
	for i := 0; i < len(bag); i++ {
		if answer == bag[i].Name {
			if answer == "Potion bas de gamme" {
				if p1.Hp != p1.Max_Hp {
					p1.Hp += bag[i].Value
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					bag = append(bag[:i], bag[i+1:]...)
					valide = true
				} else {
					fmt.Println("\033[97;3mVous avez deja toute votre vie\033[0m")
				}
			} else if answer == "Potion" {
				if p1.Hp != p1.Max_Hp {
					p1.Hp += bag[i].Value
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					bag = append(bag[:i], bag[i+1:]...)
					valide = true
				} else {
					fmt.Println("\033[97;3mVous avez deja toute votre vie\033[0m")
				}
			} else if answer == "Potion de bonne qualité" {
				if p1.Hp != p1.Max_Hp {
					p1.Hp += bag[i].Value
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					bag = append(bag[:i], bag[i+1:]...)
					valide = true
				} else {
					fmt.Println("\033[97;3mVous avez deja toute votre vie\033[0m")
				}
			} else if answer == "Potion de tres bonne qualité" {
				if p1.Hp != p1.Max_Hp {
					p1.Hp += bag[i].Value
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					bag = append(bag[:i], bag[i+1:]...)
					valide = true
				} else {
					fmt.Println("\033[3mVous avez deja toute votre vie\033[0m")
				}
			} else if answer == "Potion sacre" {
				if p1.Hp != p1.Max_Hp {
					p1.Hp += bag[i].Value
					if p1.Hp > p1.Max_Hp {
						p1.Hp = p1.Max_Hp
					}
					bag = append(bag[:i], bag[i+1:]...)
					valide = true
				} else {
					fmt.Println("\033[97;3mVous avez deja toute votre vie\033[0m")
				}
			}
		}
	}
	return p1,bag, valide
}


func Get_item(bag []structure.Item) []structure.Item {
	scanner := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())			// GET RANDOM ITEM AFTER FIGHT
	min := 1									// CAN DECIDE IF WE KEEP IT
	max := 100
	item_random := (rand.Intn(max) + min)
	if item_random <= 45 {					
		fmt.Println("\033[97;1m\033[97;4mVous avez trouvé une Potion bas de gamme\033[0m")
		fmt.Println("\033[97;4mSouhaitez vous la recuperer ?\033[0m y/n")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			bag = append(bag, structure.Item{Name: "Potion bas de gamme", Value: 10, Rarity: 1})
		} else {
			fmt.Println("\033[97;4mVous decidez de ne pas la prendre\033[0m")
		}
	} else if item_random > 45 && item_random <= 70 {
		fmt.Println("\033[97;1m\033[97;4mVous avez trouvé une Potion\033[0m")
		fmt.Println("\033[97;4mSouhaitez vous la recuperer ?\033[0m y/n")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			bag = append(bag, structure.Item{Name: "Potion", Value: 20, Rarity: 2})
		} else {
			fmt.Println("\033[97;4mVous decidez de ne pas la prendre\033[0m")
		}
	} else if item_random > 70 && item_random <= 85 {
		fmt.Println("\033[97;1m\033[97;4mVous avez trouvé une Potion de bonne qualité\033[0m")
		fmt.Println("\033[97;4mSouhaitez vous la recuperer ?\033[0m y/n")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			bag = append(bag, structure.Item{Name: "Potion de bonne qualité", Value: 30, Rarity: 3})
		} else {
			fmt.Println("\033[97;4mVous decidez de ne pas la prendre\033[0m")
		}
	} else if item_random > 85 && item_random <= 95 {
		fmt.Println("\033[97;1m\033[97;4mVous avez trouvé une Potion de très bonne qualité\033[0m")
		fmt.Println("\033[97;4mSouhaitez vous la recuperer ?\033[0m y/n")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			bag = append(bag, structure.Item{Name: "Potion de tres bonne qualité", Value: 40, Rarity: 4})
		} else {
			fmt.Println("\033[97;4mVous decidez de ne pas la prendre\033[0m")
		}
	} else {
		fmt.Println("\033[97;1m\033[97;4mVous avez trouvé une Potion sacré\033[0m")
		fmt.Println("\033[97;4mSouhaitez vous la recuperer ?\033[0m y/n")
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" || answer == "Y" {
			bag = append(bag, structure.Item{Name: "Potion sacre", Value: 50, Rarity: 5})
		} else {
			fmt.Println("\033[97;4mVous decidez de ne pas la prendre\033[0m")
		}
	}
	return bag
}
