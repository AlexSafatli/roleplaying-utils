package main

import (
	"../5etools-json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Encounter struct {
	Session int
	Level   int
	Room    int
	XP      int
}

type Participant struct {
	Name       string
	Initiative int
	HP         int
}

type EncounterData struct {
	Encounter
	Participants []Participant
}

func main() {
	if len(os.Args) < 4 {
		panic("Requires an output JSON file, input JSON file folder, and a list of monster names.")
	}

	var monsters map[string]*fiveEtoolsjson.Monster

	j, err := filepath.Glob(os.Args[2] + string(os.PathSeparator) + "*.json")
	if err != nil {
		panic(err)
	}

	monsters = make(map[string]*fiveEtoolsjson.Monster)
	for _, path := range j {
		var parsed []fiveEtoolsjson.Monster
		parsed = fiveEtoolsjson.Get5etoolsMonsters(path)
		for i := range parsed {
			if _, ok := monsters[parsed[i].Name]; !ok {
				monsters[parsed[i].Name] = &parsed[i]
			}
		}
	}

	var output EncounterData
	var monsterName string
	var monsterQty int
	for i, val := range os.Args[3:] {
		if i%2 == 0 {
			monsterName = val
		} else {
			monsterQty, err = strconv.Atoi(val)
			if err != nil {
				panic("Expected quantity for " + monsterName)
			}
			if m, ok := monsters[monsterName]; ok {
				var i int
				var p Participant
				p.Name = m.Name
				p.Initiative = (m.Dex - 10) / 2
				p.HP = m.HP.Average
				fmt.Printf("Monster '%s' (%d): %+v\n", m.Name, monsterQty, m)
				for i = 0; i < monsterQty; i++ {
					output.Participants = append(output.Participants, p)
				}
			} else {
				fmt.Printf("Could not find monster '%s'\n", monsterName)
			}
		}
	}

	file, _ := json.MarshalIndent(output, "", " ")
	if err = ioutil.WriteFile(os.Args[1], file, 0644); err != nil {
		panic(err)
	}
}
