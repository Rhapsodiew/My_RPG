package main

import "fmt"

type Character struct {
	Name       string
	Health     int
	Max_health int
	Strength   int
}

func main() {
	link := Character{Name: "Link", Health: 60, Max_health: 60, Strength: 15}

	goblin := Character{Name: "Bokoblin", Health: 30, Max_health: 30, Strength: 5}
	ganon := Character{Name: "Ganon", Health: 150, Max_health: 150, Strength: 20}

	enemy := []Character{
		goblin,
		ganon,
	}

	etage := 1

	fmt.Println("Navi : Link, c'est ici que ce trouve Ganon.")
	var text_pass rune
	fmt.Scan(&text_pass)
	fmt.Println("Navi : Nous devons grimper tout en haut de la tour et le vaincre pour liberer Hyrule du fleau qu'est Ganon.")
	fmt.Scan(&text_pass)
	fmt.Println("Navi : Allons - y !")
	fmt.Scan(&text_pass)
	fmt.Println()
	fmt.Println("========== CHATEAU D'HYRULE ==========")
	fmt.Println()
	fmt.Println("========== Etage 1 ==========")
	fmt.Println()
	fmt.Println("Navi : Nous voici au premier étage.")
	fmt.Scan(&text_pass)
	fmt.Println("Navi : Regarde ! Il y a un Bokoblin la-bas. Va-t'en debarasser. ")
	fmt.Scan(&text_pass)

	if etage == 1 {
		fight1(link, enemy[0])
		etage++
	}
	fmt.Println("Navi : Heureusement que ce Bokoblin etait déja blessé.")
	fmt.Scan(&text_pass)
	fmt.Println("Navi : Restons sur nos garde.")
	fmt.Scan(&text_pass)

	for etage > 1 && etage < 10{
		fmt.Println()
		fmt.Printf("========== Etage %d ==========\n",etage)
		fmt.Println()
		result := fight(&link,enemy[0])
		if result == false{
			fmt.Println("Vous avez echouer !")
			return
		}else{
			fmt.Println("Navi : Bien jouer, passons au prochaine etage !")
			etage++
		}
	}
	if etage == 10{
		fmt.Println()
		fmt.Printf("========== Etage %d ==========\n",etage)
		fmt.Println()
		fmt.Println("Ganon : Vous voila. Je t'attendais vermine !")
		fmt.Scan(&text_pass)
		fmt.Println("Navi : C'est donc lui 'Ganon le Fleau'.")
		fmt.Scan(&text_pass)
		fmt.Println("Navi : Reste sur tes gardes.")
		fmt.Scan(&text_pass)
		fmt.Println("Ganon : Vous pensez pouvoir m'arreter ?")
		fmt.Scan(&text_pass)
		fmt.Println("Ganon : AHAHAHAHAH")
		fmt.Println("Ganon : Que vous êtes bêtes !")
		fmt.Scan(&text_pass)
		fmt.Println("Ganon : JE SUIS GANON LE FLEAU !!")
		fmt.Println("Ganon : ET RIEN NE POURRA ME STOPPER DANS MA CONQUETE D'HYRULE !!")
		fmt.Scan(&text_pass)
		fmt.Println()
		result := fight_boss(&link,ganon)
		if result == false {
			fmt.Println("Vous avez echouer !")
			return
		}else{
			fmt.Println("Navi : Tu as sauver Hyrule de ce monstre !")
			fmt.Println("Navi : Hyrule pourra enfin connaitre la paix.")
		}
	}
}
