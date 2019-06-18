package main

import (
	"../5etools-json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	if len(os.Args) >= 3 {
		if os.Args[1] == "spell" {
			var spells []fiveEtoolsjson.Spell
			for _, path := range os.Args[2:] {
				pathSpells := fiveEtoolsjson.Get5etoolsSpells(path)
				spells = append(spells, pathSpells...)
			}
			i := rand.Intn(len(spells))
			randomSpell := spells[i]
			fmt.Printf("%d %s %s\n%s\n", i, randomSpell.Name, randomSpell.Source, randomSpell.Entries)
		}
	}
}
