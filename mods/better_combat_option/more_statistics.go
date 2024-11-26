package bettercombatoption

import (
	"fmt"
	"math/rand"
	"mods/structure"
	"time"
)

func Is_faster(p1 *structure.Character, p2 *structure.Character) bool {
	if p1.Spd >= p2.Spd {
		return true
	} else {
		return false
	}
}

func Dodge(p1 *structure.Character, p2 *structure.Character) bool {
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())
	dodge := (rand.Intn(max) + min)
	if dodge <= p1.Spd-p2.Spd {
		fmt.Printf("\033[97;1m%s\033[0m\033[97;3m esquive l'attaque de \033[0m\033[97;1m%s\033[0m",p1.Name,p2.Name)
		return true
	}
	return false
}

func Crit(p1 *structure.Character,dmg float64) float64 {
	min := 1
	max := 100
	rand.Seed(time.Now().UnixNano())
	crit1 := (rand.Intn(max) + min)
	if crit1 <= p1.Luck {
		fmt.Printf("\033[97;1m%v\033[0m\033[97;3m a fait un \033[97;1mCOUP CRITIQUE\033[0m\n",p1.Name)
		return dmg*2
	}
	return dmg
}

func Calculate_type_attk(p1 *structure.Character) float64 {
	var dmg1 float64
	classes := structure.Read_classes()
	for _, class := range classes {
		if p1.Class == class.Id {
			if class.Attack_type == "physical" {
				dmg1 = float64(p1.Str)
			} else if class.Attack_type == "magical" {
				dmg1 = float64(p1.Int)
			}
		}
	}
	return dmg1
}

func Calculate_def_res(p1 *structure.Character, p2 *structure.Character,dmg1 float64) float64 {
	classes := structure.Read_classes()
	//================================================= DMG FOR P1
	for _, class := range classes {
		if p1.Class == class.Id {
			if class.Attack_type == "physical" {
				dmg1 = dmg1 - dmg1*(float64(p2.Def)/100)
			} else if class.Attack_type == "magical" {
				dmg1 = dmg1 - dmg1*(float64(p2.Res)/100)
			}
		}
	}
	return dmg1
}

// //================================================== DMG FOR P2
// for _, class := range classes {
// 	if p2.Class == class.Id {
// 		if class.Attack_type == "physical" {
// 			dmg2 = float64(p2.Str)
// 		} else if class.Attack_type == "magical" {
// 			dmg2 = float64(p2.Int)
// 		}
// 	}
// }

// //COUP CRIT
// rand.Seed(time.Now().UnixNano())
// crit2 := (rand.Intn(max-min+1) + min)
// if crit2 <= p2.Luck {
// 	dmg2 = dmg2 * 2
// 	fmt.Println(p2.Name, " a fait un COUP CRITIQUE")
// 	fmt.Println()
// }
// for _, class := range classes {
// 	if p2.Class == class.Id {
// 		if class.Attack_type == "physical" {
// 			dmg2 = dmg2 - dmg2*(float64(p1.Def)/100)
// 		} else if class.Attack_type == "magical" {
// 			dmg2 = dmg2 - dmg2*(float64(p1.Res)/100)
// 		}
// 	}
// }

// if p1.Spd >= p2.Spd {
// 	rand.Seed(time.Now().UnixNano())
// 	dodge := (rand.Intn(max-min+1) + min)
// 	if dodge <= p1.Spd-p2.Spd {
// 		fmt.Println("Vous avez esquivé l'attaque")
// 		dmg2 = 0
// 	}
// 	return p1.Name, dmg1, dmg2
// } else {
// 	rand.Seed(time.Now().UnixNano())
// 	dodge := (rand.Intn(max-min+1) + min)
// 	if dodge <= p2.Spd-p1.Spd {
// 		fmt.Println(p2.Name, " a esquivé l'attaque")
// 		dmg1 = 0
// 	}
// 	return p2.Name, dmg2, dmg1
// }
// }
