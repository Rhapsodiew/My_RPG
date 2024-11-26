package characteristics

import (
	"fmt"
	"mods/structure"
)

func Option_character() {
	fmt.Print("\033[97;1mC. Character\033[0m		")
	fmt.Println()	
}

func Show_stats(p1 *structure.Character) {
	classes := structure.Read_classes()
	races := structure.Read_races()
	
	fmt.Println("\033[97;1mName :\033[0m ", p1.Name)
	fmt.Println("\033[97;1mHP :\033[0m ", p1.Hp, "/", p1.Max_Hp)
	fmt.Println("\033[97;1mMp :\033[0m ", p1.Mp)
	fmt.Println("\033[97;1mStr :\033[0m ", p1.Str)
	fmt.Println("\033[97;1mInt :\033[0m ", p1.Int)
	fmt.Println("\033[97;1mDef :\033[0m ", p1.Def)
	fmt.Println("\033[97;1mRes :\033[0m ", p1.Res)
	fmt.Println("\033[97;1mSpd :\033[0m ", p1.Spd)
	fmt.Println("\033[97;1mLuck :\033[0m ", p1.Luck)
	for _ , race := range races{
		if p1.Race == race.Id{
			fmt.Println("\033[97;1mRace :\033[0m ", race.Name)
		}
	}
	for _ , class := range classes{
		if p1.Class == class.Id{
			fmt.Println("\033[97;1mClass :\033[0m ", class.Name)
		}
	}
}

func Is_Effective(base_dmg,dmg float64, p1 *structure.Character) {
	if base_dmg != dmg {
		if base_dmg < dmg {
			fmt.Println("\033[97;3mL'attaque de \033[0m\033[97;1m", p1.Name, "\033[0m\033[97;3m est \033[0m\033[97;1m\033[97;4msuper efficace\033[0m")
			
		} else {
			fmt.Println("\033[97;3mL'attaque de \033[0m\033[97;1m", p1.Name, "\033[0m\033[97;3m n'est \033[0m\033[97;1m\033[97;4mpas tres efficace\033[0m")
			
		}
	}
}

func Calculate_race_class(p1 *structure.Character, p2 *structure.Character, dmg1 float64) float64 {
	races := structure.Read_races()
	classes := structure.Read_classes()
	//CALCULATE RACE
	for _, race := range races {
		if p1.Race == race.Id {
			list_strength_race := race.Strength
			list_weakness_race := race.Weakness
			for _, strength_race := range list_strength_race {
				if strength_race == p2.Race {
					dmg1 = dmg1 * 2
				}
			}
			for _, weakness_race := range list_weakness_race {
				if weakness_race == p2.Race {
					dmg1 = dmg1 / 2
				}
			}
		}
	}
	//CALCULATE CLASS
	for _, class := range classes {
		if p1.Class == class.Id {
			list_strength_class := class.Strengths
			list_weakness_class := class.Weaknesses
			for _, strength_class := range list_strength_class {
				if strength_class == p2.Class {
					dmg1 = dmg1 * 2
				}
			}
			for _, weakness_class := range list_weakness_class {
				if weakness_class == p2.Class {
					dmg1 = dmg1 / 2
				}
			}
		}
	}
	return dmg1
}
	// //P2
	// for _, race := range races {
	// 	if p2.Race == race.Id {
	// 		list_strength_race := race.Strength
	// 		list_weakness_race := race.Weakness
	// 		for _, strength_race := range list_strength_race {
	// 			if strength_race == p1.Race {
	// 				dmg2 = dmg2 * 2
	// 			}
	// 		}
	// 		for _, weakness_race := range list_weakness_race {
	// 			if weakness_race == p1.Race {
	// 				dmg2 = dmg2 / 2
	// 			}
	// 		}
	// 	}
	// }
	// for _, class := range classes {
	// 	if p2.Class == class.Id {
	// 		list_strength_class := class.Strengths
	// 		list_weakness_class := class.Weaknesses
	// 		for _, strength_class := range list_strength_class {
	// 			if strength_class == p1.Class {
	// 				dmg2 = dmg2 * 2
	// 			}
	// 		}
	// 		for _, weakness_class := range list_weakness_class {
	// 			if weakness_class == p1.Class {
	// 				dmg2 = dmg2 / 2
	// 			}
	// 		}
	// 	}
	// }

	// if base_dmg2 != dmg2 {
	// 	if base_dmg2 < dmg2 {
	// 		fmt.Println("L'attaque de ", p2.Name, " est super efficace contre ", p1.Name)
	// 	} else {
	// 		fmt.Println("L'attaque de ", p2.Name, " n'est pas tres efficace contre ", p1.Name)
	// 	}
	// }

	// return dmg1, dmg2

