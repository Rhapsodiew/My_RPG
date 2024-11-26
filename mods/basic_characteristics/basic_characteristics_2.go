package characteristics

import (
	"bufio"
	"fmt"
	"mods/structure"
	"os"
)

// SHOW ALIGNMENT IN C MENU
func Show_stats_2 (p1 *structure.Character){
	classes := structure.Read_classes()
	for _, class := range classes{
		if p1.Class == class.Id{
			fmt.Printf("\033[97;1mAlignment :\033[0m %v\n",class.Alignment)
		}
	}
}
// NEW OPTION
func Option_special_move(){
	fmt.Println("\033[97;1mSp. Special\033[0m")
}

func Special_move(used bool ,enemy *structure.Character) bool{
	scanner := bufio.NewScanner(os.Stdin)

	if used{
		fmt.Println("\033[97;4mVous ne pouvez faire cela qu'un fois par combat\033[97;0m")
	}else{
		fmt.Println("\033[97;4m\033[97;1mSPECIAL ATTACK !!\033[0m")
		fmt.Println()
		fmt.Println("\033[97;4mVous embraser l'ennemi et l'attaquer sans qu'il ne puisse rien faire !")
		scanner.Scan()
		scanner.Text()
		enemy.Hp = enemy.Hp /5
	}
	return true
}