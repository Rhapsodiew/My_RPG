package main

import (
	"bufio"
	"fmt"
	"mods/structure"
	"os"
	"strings"
)

func Introduction(p1, boss structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\033[36;1mNavi \033[0m: On est enfin arrivé au pied du CHATEAU D'HYRULE.")
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[36;1mNavi \033[0m: %s, c'est ici que ce trouve %s.\n", p1.Name, boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[36;1mNavi \033[0m: Nous devons grimper tout en haut du chateau et le vaincre pour liberer Hyrule du fleau qu'est %s.\n", boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Println("\033[36;1mNavi \033[0m: Allons - y !")
	scanner.Scan()
	scanner.Text()
	fmt.Println()
	fmt.Println("\033[33;1m \033[33;4m========== CHATEAU D'HYRULE ==========\033[0m")
	fmt.Println()
}

func Boss_intro(boss structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\033[31;1m%s \033[0m: Vous voila. Je t'attendais vermine !\n", boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[36;1mNavi \033[0m: C'est donc lui '\033[31;1m%s\033[0m'.",boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Println("\033[36;1mNavi \033[0m: Reste sur tes gardes.")
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s \033[0m: Vous pensez pouvoir m'arreter ?\n", boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s \033[0m: AHAHAHAHAH\n", boss.Name)
	fmt.Printf("\033[31;1m%s \033[0m: Que vous êtes bêtes !\n", boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s \033[0m: JE SUIS %s !!\n", boss.Name, strings.ToUpper(boss.Name))
	fmt.Printf("\033[31;1m%s \033[0m: ET RIEN NE POURRA ME STOPPER DANS MA CONQUETE D'HYRULE !!\n", boss.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Println()
}

func Boss_phase2(enemy structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\033[31;1m%s\033[0m : Tu n'est donc pas arriver ici pour rien.\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : \033[97;1m\033[97;3mRWHAAAA\033[0m\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Println("\033[36;1mNavi \033[0m: Oh non, son corp se regenere et devient encore plus grand qu'avant")
	scanner.Scan()
	scanner.Text()
	fmt.Println()
	enemy.Hp += enemy.Max_Hp / 4
	enemy.Str += 5
	enemy.Int += 5
	enemy.Def += 8
	enemy.Res += 8
	enemy.Spd -= 2
}

func Boss_mid_end(enemy structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\033[31;1m%s\033[0m : Keugh", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Tu sais ...", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Il te reste encore du chemin pour arriver au sommet", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Ne te rejouit pas si vite de ta victoire", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[97;3mVous avez vaincu \033[97;1m%s\033[0m", enemy.Name)
	scanner.Scan()
	scanner.Text()
}

func Boss_u_die(enemy structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\033[31;1m%s\033[0m : Idiot, tu pensais être de taille face a MOI, GANON, AHAHAHAH !!!!\n", enemy.Name)
	fmt.Println()
	scanner.Scan()
	scanner.Text()
}
func Boss_end(enemy structure.Character) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\033[31;1m%s\033[0m : Keugh\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : ...\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Ne t'inquiete pas ...\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Je reviendrais plus fort que jamais ...\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[31;1m%s\033[0m : Et Hyrule sombrera dans la terreur !\n", enemy.Name)
	scanner.Scan()
	scanner.Text()
	fmt.Printf("\033[97;3mVous avez vaincu \033[97;1m%s\033[0m", enemy.Name)
	scanner.Scan()
	scanner.Text()
}

func Etage_3(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[36;1mNavi \033[0m: Les ennemis vont surement devenir plus coriace en montant d'étage. Fait attention.")
	scanner.Scan()
	scanner.Text()
}

func Etage_mid(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\033[36;1mNavi \033[0m: On en est qu'a la moitié ...")
	scanner.Scan()
	scanner.Text()
	fmt.Println("\033[36;1mNavi \033[0m: Ne baisse pas les bras. Tu dois sauver HYRULE !")
}

func Traps(traps structure.Traps){
	scanner := bufio.NewScanner(os.Stdin)

	if traps.Name == "Treasure Room"{
		fmt.Println("\033[36;1mNavi \033[0m: Dommage qu'il n'y avait pas plus a prendre. On a eu que quelque piece ...")
	}else if traps.Name == "Pike Wall"{
		fmt.Println("\033[36;1mNavi \033[0m: Pourquoi il y a autant de piege ? C'est vraiment fatigant")
	}else if traps.Name == "Secret Code"{
		fmt.Println("\033[36;1mNavi \033[0m: Pourquoi il y a autant de piege ? Je pensais que les monstres etaient betes ...")
	}
	scanner.Scan()
	scanner.Text()
}