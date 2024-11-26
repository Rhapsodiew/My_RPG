package main

import (
	"mods/structure"
)
func change_mob_stats(difficulty int) ([]structure.Character, []structure.Character) {
	enemies := structure.Read_enemies()
	bosses := structure.Read_bosses()
	if difficulty == 2 {
		for i, mobs := range enemies {

			enemies[i].Hp = mobs.Hp * 3 / 2
			enemies[i].Mp = mobs.Mp * 3 / 2
			enemies[i].Str = mobs.Str * 3 / 2
			enemies[i].Int = mobs.Int * 3 / 2
			enemies[i].Def = mobs.Def * 3 / 2
			enemies[i].Res = mobs.Res * 3 / 2
			enemies[i].Spd = mobs.Spd * 3 / 2
			enemies[i].Luck = mobs.Luck * 3 / 2

		}
		for i, mobs := range bosses {

			bosses[i].Hp = mobs.Hp * 3 / 2
			bosses[i].Mp = mobs.Mp * 3 / 2
			bosses[i].Str = mobs.Str * 3 / 2
			bosses[i].Int = mobs.Int * 3 / 2
			bosses[i].Def = mobs.Def * 3 / 2
			bosses[i].Res = mobs.Res * 3 / 2
			bosses[i].Spd = mobs.Spd * 3 / 2
			bosses[i].Luck = mobs.Luck * 3 / 2

		}
	} else if difficulty == 3 {
		for i, mobs := range enemies {

			enemies[i].Hp = mobs.Hp * 2
			enemies[i].Mp = mobs.Mp * 2
			enemies[i].Str = mobs.Str * 2
			enemies[i].Int = mobs.Int * 2
			enemies[i].Def = mobs.Def * 2
			enemies[i].Res = mobs.Res * 2
			enemies[i].Spd = mobs.Spd * 2
			enemies[i].Luck = mobs.Luck * 2

		}
		for i, mobs := range bosses {

			bosses[i].Hp = mobs.Hp * 2
			bosses[i].Mp = mobs.Mp * 2
			bosses[i].Str = mobs.Str * 2
			bosses[i].Int = mobs.Int * 2
			bosses[i].Def = mobs.Def * 2
			bosses[i].Res = mobs.Res * 2
			bosses[i].Spd = mobs.Spd * 2
			bosses[i].Luck = mobs.Luck * 2
		}
	}
	return enemies, bosses
}
