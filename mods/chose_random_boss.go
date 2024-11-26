package main

import (
	// "fmt"
	"math/rand"
	"time"
	"mods/structure"
)

func chose_random_boss(bosses []structure.Character) structure.Character{
	

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	boss_random := (rand.Intn(max) + min)
	// fmt.Println(boss_random,"boss random")
	//CHOSE RANDOM BOSS					HIDE PRINT BOSS !!
	var boss structure.Character
	var chose_boss []structure.Character
	if boss_random <= 45 {
		// fmt.Println("1")
		for i := 0; i < len(bosses); i++ {
			if bosses[i].Rarity == 1 {
				chose_boss = append(chose_boss, bosses[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = len(chose_boss)
		boss_random = (rand.Intn(max) + min)
		boss = chose_boss[boss_random-1]
		// fmt.Println(boss)
	} else if boss_random > 45 && boss_random <= 70 {
		// fmt.Println("2")
		for i := 0; i < len(bosses); i++ {
			if bosses[i].Rarity == 2 {
				chose_boss = append(chose_boss, bosses[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = len(chose_boss)
		boss_random = (rand.Intn(max) + min)
		boss = chose_boss[boss_random-1]
		// fmt.Println(boss)
	} else if boss_random > 70 && boss_random <= 85 {
		// fmt.Println("3")
		for i := 0; i < len(bosses); i++ {
			if bosses[i].Rarity == 3 {
				chose_boss = append(chose_boss, bosses[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = len(chose_boss)
		boss_random = (rand.Intn(max) + min)
		boss = chose_boss[boss_random-1]
		// fmt.Println(boss)
	} else if boss_random > 85 && boss_random <= 95 {
		// fmt.Println("4")
		for i := 0; i < len(bosses); i++ {
			if bosses[i].Rarity == 4 {
				chose_boss = append(chose_boss, bosses[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = len(chose_boss)
		boss_random = (rand.Intn(max) + min)
		boss = chose_boss[boss_random-1]
		// fmt.Println(boss)
	} else {
		// fmt.Println("5")
		for i := 0; i < len(bosses); i++ {
			if bosses[i].Rarity == 5 {
				chose_boss = append(chose_boss, bosses[i])
			}
		}
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = len(chose_boss)
		boss_random = (rand.Intn(max) + min)
		boss = chose_boss[boss_random-1]
		// fmt.Println(boss)
	}
	return boss
}
