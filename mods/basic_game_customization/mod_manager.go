package gamecustom

import (
	"bufio"
	"fmt"
	"os"
)

func Option_mod() {
	fmt.Println("\033[97;1mM. 	Mod Manager\033[0m")
}

func Is_mod_activate(mods []string, name string) bool {
	for _, mod_in := range mods {
		if name == mod_in {
			return true
		}
	}
	return false
}

func Add_mods(mods []string, to_add string) []string {
	for i, mod_in := range mods {
		if mod_in == to_add {
			fmt.Printf("%s est déja activé\n", to_add)
			mods = append(mods[:i], mods[i+1:]...)
		}
	}
	mods = append(mods, to_add)
	return mods
}

func Check_mods(start string, mods []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	if start == "M" || start == "Mods Manager" || start == "Mods" {
		var answer string
		for answer != "Q" {
			fmt.Println()
			fmt.Println("\033[33;1m-> \033[97;4mChoisissez les mods a activer\033[0m")
			fmt.Println()
			fmt.Println("\033[97;1mQ. Quitter\033[0m")
			fmt.Println()
			fmt.Println("\033[97;4mListe des mods :\033[0m")
			fmt.Println()
			fmt.Println("A. Activate All Mods")
			fmt.Println("1. Better Combat Option")
			fmt.Println("2. Basic Characteristics")
			fmt.Println("3. Basic Game Customization")
			scanner.Scan()
			answer = scanner.Text()
			fmt.Println()
			if answer == "A" {
				mods = Add_mods(mods, "Better_Combat_Option")
				mods = Add_mods(mods, "More_Statistics")
				mods = Add_mods(mods, "Inventory")
				mods = Add_mods(mods, "Magic_Skills")
				mods = Add_mods(mods, "Basic_Characteristics")
				mods = Add_mods(mods, "Basic_Characteristics_2")
				mods = Add_mods(mods, "Level_And_Experience")
				mods = Add_mods(mods, "Random_Game_Events")
			}
			if answer == "1" {
				fmt.Println("1. Better_Combat_Option")
				fmt.Println("2. More_Statistics (Need Better_Combat_Option)")
				fmt.Println("3. Inventory (Need Better_Combat_Option)")
				fmt.Println("4. Magic_Skills (Need Better_Combat_Option)")
				scanner.Scan()
				mod_to_add := scanner.Text()
				//=================== BETTER COMBAT OPTION ======================
				if mod_to_add == "1" {
					for i, mod_in := range mods {
						if mod_in == "Better_Combat_Option" {
							fmt.Println("Better_Combat_Option est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Better_Combat_Option")

				} else if mod_to_add == "2" {
					for i, mod_in := range mods {
						if mod_in == "Better_Combat_Option" {
							fmt.Println("Better_Combat_Option est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Better_Combat_Option")

					for i, mod_in := range mods {
						if mod_in == "More_Statistics" {
							fmt.Println("More_Statistics est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "More_Statistics")

				} else if mod_to_add == "3" {
					for i, mod_in := range mods {
						if mod_in == "Better_Combat_Option" {
							fmt.Println("Better_Combat_Option est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Better_Combat_Option")

					for i, mod_in := range mods {
						if mod_in == "Inventory" {
							fmt.Println("Inventory est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Inventory")

				} else if mod_to_add == "4" {
					for i, mod_in := range mods {
						if mod_in == "Better_Combat_Option" {
							fmt.Println("Better_Combat_Option est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Better_Combat_Option")

					for i, mod_in := range mods {
						if mod_in == "Magic_Skills" {
							fmt.Println("Magic_Skills est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Magic_Skills")
				}

				//================ BASIC CHARACTERISTICS ===================

			} else if answer == "2" {
				fmt.Println("1. Basics_Characteristics")
				fmt.Println("2. Basic_Characteristics_2 (Need Basics_Characteristics)")
				fmt.Println("3. Level_And_Experience (Need Basics_Characteristics)")
				scanner.Scan()
				mod_to_add := scanner.Text()

				if mod_to_add == "1" {
					for i, mod_in := range mods {
						if mod_in == "Basic_Characteristics" {
							fmt.Println("Basic_Characteristics est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Basic_Characteristics")

				} else if mod_to_add == "2" {
					for i, mod_in := range mods {
						if mod_in == "Basic_Characteristics" {
							fmt.Println("Basic_Characteristics est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Basic_Characteristics")

					for i, mod_in := range mods {
						if mod_in == "Basic_Characteristics_2" {
							fmt.Println("Basic_Characteristics_2 est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Basic_Characteristics_2")

				} else if mod_to_add == "3" {
					for i, mod_in := range mods {
						if mod_in == "Basic_Characteristics" {
							fmt.Println("Basic_Characteristics est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Basic_Characteristics")

					for i, mod_in := range mods {
						if mod_in == "Level_And_Experience" {
							fmt.Println("Level_And_Experience est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Level_And_Experience")
				}

				//===================== BASIC GAME CUSTOMIZATION ==============================
			} else if answer == "3" {
				fmt.Println("1. Random_Game_Events")
				scanner.Scan()
				mod_to_add := scanner.Text()

				if mod_to_add == "1" {
					for i, mod_in := range mods {
						if mod_in == "Random_Game_Events" {
							fmt.Println("Random_Game_Events est déja activé")
							mods = append(mods[:i], mods[i+1:]...)
						}
					}
					mods = append(mods, "Random_Game_Events")
				}
			}
		}
	}
	return mods
}
