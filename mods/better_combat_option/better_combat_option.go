package bettercombatoption

import (
	"fmt"
	"mods/structure"
)
func Better_combat_option() {
	fmt.Println("\033[97;1mD. Defend\033[0m 		\033[97;1mE. Escape\033[0m")
	fmt.Println()
}

// USELESS FUNC CAN DELETE
func Chose_option(end_turn *bool,p1 *structure.Character, enemy *structure.Character,action string) (*structure.Character, structure.Character,bool,string) {
	fmt.Println()
	switch action {
	case "A", "Attack":
		fmt.Println()
		*end_turn = true
	case "S", "Skills":
		*end_turn = true
	case "D", "Defend":
		fmt.Println("Vous vous mettez en garde !")
		fmt.Println()
		*end_turn = true
	case "E", "Escape":
		fmt.Printf("En essayant de fuir, vous subissez des attaques de %v\n", enemy.Name)
		fmt.Println("\033[36;1mNavi \033[0m: Tu n'as pas honte de fuir ?")
		*end_turn = true
	default:
		fmt.Println("\033[36;1mNavi \033[0m: Qu'est ce que tu essayes de faire ?")
		fmt.Println()
	}
	return p1, *enemy, *end_turn,action
}

func Defend (dmg float64) float64{
	dmg = dmg /2
	return dmg
}
