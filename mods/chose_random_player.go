package main

import (

	"math/rand"
	"mods/structure"
	"time"
)

func chose_random_player() structure.Character {
	players := structure.Read_players()
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	chara_random := (rand.Intn(max) + min)
	//CHOSE RANDOM PLAYER 					HIDE PRINT PLAYER
	var p1 structure.Character
	if chara_random <= 45 {
		p1 = players[0]
		p1.Max_Hp = p1.Hp
		p1.Max_Mp = p1.Mp
		// fmt.Println(p1)
	} else if chara_random > 45 && chara_random <= 70 {
		p1 = players[1]
		p1.Max_Hp = p1.Hp
		p1.Max_Mp = p1.Mp
		// fmt.Println(p1)
	} else if chara_random > 70 && chara_random <= 85 {
		p1 = players[2]
		p1.Max_Hp = p1.Hp
		p1.Max_Mp = p1.Mp
		// fmt.Println(p1)
	} else if chara_random > 85 && chara_random <= 95 {
		p1 = players[3]
		p1.Max_Hp = p1.Hp
		p1.Max_Mp = p1.Mp
		// fmt.Println(p1)
	} else {
		p1 = players[4]
		p1.Max_Hp = p1.Hp
		p1.Max_Mp = p1.Mp
		// fmt.Println(p1)
	}
	// fmt.Println(p1,"p1")
	return p1
}
