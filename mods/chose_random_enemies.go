package main

import (
	"math/rand"
	"time"
	"mods/structure"
)

func chose_random_enemies(etage int, enemies []structure.Character) structure.Character {
	var chose_enemies []structure.Character
	var stage_enemies structure.Character
	if etage%10 != 0 {
		min := 1
		max := 100
		//=========== CHOSE RANDOM ENEMIES ==============
		rand.Seed(time.Now().UnixNano())
		min = 1
		max = 100
		enemies_random := (rand.Intn(max) + min)

		// fmt.Println(enemies_random,"enemie random")

		if enemies_random <= 45 {
			// fmt.Println("1")

			for i := 0; i < len(enemies); i++ {
				if enemies[i].Rarity == 1 {
					chose_enemies = append(chose_enemies, enemies[i])
				}
			}
			// fmt.Println(chose_enemies)
			rand.Seed(time.Now().UnixNano())
			min = 1

			enemies_random = (rand.Intn(len(chose_enemies)) + min)
			stage_enemies = chose_enemies[enemies_random-1]
			// fmt.Println(stage_enemies)
		} else if enemies_random > 45 && enemies_random <= 70 {
			// fmt.Println("2")
			for i := 0; i < len(enemies); i++ {
				if enemies[i].Rarity == 2 {
					chose_enemies = append(chose_enemies, enemies[i])
				}
			}
			rand.Seed(time.Now().UnixNano())
			min = 1

			enemies_random = (rand.Intn(len(chose_enemies)) + min)
			stage_enemies = chose_enemies[enemies_random-1]
			// fmt.Println(stage_enemies)
		} else if enemies_random > 70 && enemies_random <= 85 {
			// fmt.Println("3")
			for i := 0; i < len(enemies); i++ {
				if enemies[i].Rarity == 3 {
					chose_enemies = append(chose_enemies, enemies[i])
				}
			}
			rand.Seed(time.Now().UnixNano())
			min = 1
			enemies_random = (rand.Intn(len(chose_enemies)) + min)
			stage_enemies = chose_enemies[enemies_random-1]
			// fmt.Println(stage_enemies)
		} else if enemies_random > 85 && enemies_random <= 95 {
			// fmt.Println("4")
			for i := 0; i < len(enemies); i++ {
				if enemies[i].Rarity == 4 {
					chose_enemies = append(chose_enemies, enemies[i])
				}
			}
			rand.Seed(time.Now().UnixNano())
			min = 1
			enemies_random = (rand.Intn(len(chose_enemies)) + min)
			stage_enemies = chose_enemies[enemies_random-1]
			// fmt.Println(stage_enemies)
		} else {
			// fmt.Println("5")
			for i := 0; i < len(enemies); i++ {
				if enemies[i].Rarity == 5 {
					chose_enemies = append(chose_enemies, enemies[i])
				}
			}
			rand.Seed(time.Now().UnixNano())
			min = 1
			enemies_random = (rand.Intn(len(chose_enemies)) + min)
			stage_enemies = chose_enemies[enemies_random-1]
			// fmt.Println(stage_enemies)
		}
	}
	return stage_enemies
}
