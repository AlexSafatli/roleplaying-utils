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
		} else if os.Args[1] == "item" {
			var items []fiveEtoolsjson.Item
			for _, path := range os.Args[2:] {
				items = append(items, fiveEtoolsjson.Get5etoolsItems(path)...)
			}
			r := items[rand.Intn(len(items))]
			fmt.Printf("%d gp %d lb %s %s\n%s\n", r.Value, r.Weight, r.Name, r.Source, r.Entries)
		}
	}
}
