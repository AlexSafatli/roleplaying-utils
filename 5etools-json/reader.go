package fiveEtoolsjson

import (
	"encoding/json"
	"io/ioutil"
)

type SpellList struct {
	Spells []Spell `json:"spell"`
}

type MonsterList struct {
	Monsters []Monster `json:"monster"`
}

type ItemList struct {
	Items []Item `json:"item"`
}

func Get5etoolsSpells(path string) []Spell {
	var spellList SpellList
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dat, &spellList)
	if err != nil {
		panic(err)
	}
	return spellList.Spells
}

func Get5etoolsMonsters(path string) []Monster {
	var monsterList MonsterList
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dat, &monsterList)
	if err != nil {
		panic(err)
	}
	return monsterList.Monsters
}

func Get5etoolsItems(path string) []Item {
	var itemList ItemList
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dat, &itemList)
	if err != nil {
		panic(err)
	}
	return itemList.Items
}
