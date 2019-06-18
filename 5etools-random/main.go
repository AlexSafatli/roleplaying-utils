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
				spells = append(spells, fiveEtoolsjson.Get5etoolsSpells(path)...)
			}
			r := spells[rand.Intn(len(spells))]
			fmt.Printf("%d %s %s\n%s\n", r.Level, r.Name, r.Source, r.Entries)
		}
	}
}
