package structure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Character struct {
	Name   string `json:"name"`
	Hp     int    `json:"hp"`
	Max_Hp int
	Mp     int `json:"mp"`
	Max_Mp int
	Str    int `json:"str"`
	Int    int `json:"int"`
	Def    int `json:"def"`
	Res    int `json:"res"`
	Spd    int `json:"spd"`
	Luck   int `json:"luck"`
	Race   int `json:"race"`
	Class  int `json:"class"`
	Rarity int `json:"rarity"`
}

type Races struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Strength []int  `json:"strength"`
	Weakness []int  `json:"weakness"`
	Rarity   string `json:"rarity"`
}

type Spells struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cost     int `json:"cost"`
	Dmg      int `json:"dmg"`
	Effect   string `json:"effect"`
	Cooldown int `json:"cooldown"`
	Race     string `json:"race"`
	Class    string `json:"class"`
	Rarity   int `json:"rarity"`
}

type Classes struct{
	Id int	`json:"id"`
	Name string	`json:"name"`
	Strengths []int	`json:"strengths"`
	Weaknesses []int	`json:"weaknesses"`
	Attack_type string	`json:"attack_type"`
	Alignment string	`json:"alignment"`
	Rarity int	`json:"rarity"`
}

type Item struct{
	Name string
	Value int
	Rarity int
}

type Traps struct{
	Id int	`json:"id"`
	Name string	`json:"name"`
	Requirement string	`json:"requirement"`
	Rarity int	`json:"rarity"`
}

func Read_bosses() []Character {
	var bosses []Character
	content_bosses, err := ioutil.ReadFile("../json/bosses.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_bosses, &bosses)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return bosses
}

func Read_players() []Character {
	var players []Character
	content_players, err := ioutil.ReadFile("../json/players.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_players, &players)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return players
}

func Read_enemies() []Character {
	var enemies []Character
	content_enemies, err := ioutil.ReadFile("../json/enemies.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_enemies, &enemies)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return enemies
}

func Read_races() []Races {
	var races []Races
	content_races, err := ioutil.ReadFile("../json/races.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_races, &races)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return races
}

func Read_spells() []Spells {
	var spells []Spells
	content_spells, err := ioutil.ReadFile("../json/spells.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_spells, &spells)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return spells
}

func Read_classes () []Classes{
	var classes []Classes
	content_classes, err := ioutil.ReadFile("../json/classes.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_classes, &classes)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return classes
}


func Read_traps() []Traps {
	var traps []Traps
	content_traps, err := ioutil.ReadFile("../json/traps.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(content_traps, &traps)
	if err != nil {
		fmt.Println("Error JSON Unmarshal")
		fmt.Println(err.Error())
	}
	return traps
}

func Read_all() ([]Character,[]Character,[]Character,[]Races,[]Spells) {
	enemies := Read_enemies()
	bosses := Read_bosses()
	player := Read_players()
	races := Read_races()
	spells := Read_spells()
	return enemies, bosses, player, races, spells
}