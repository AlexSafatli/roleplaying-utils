package fiveEtoolsjson

import (
	"encoding/json"
	"io/ioutil"
)

type SpellList struct {
	Spells []Spell `json:"spell"`
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
