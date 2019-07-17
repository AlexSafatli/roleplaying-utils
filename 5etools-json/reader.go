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
