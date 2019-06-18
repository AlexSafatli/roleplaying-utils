package fiveEtoolsjson

type Spell struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	School string `json:"school"`
	//Time SpellTime `json:"time"`
	Range       SpellRange    `json:"range"`
	Classes     SpellClasses  `json:"classes"`
	Source      string        `json:"source"`
	Entries     []interface{} `json:"entries"`
	Page        int           `json:"page"`
	DamageTypes []string      `json:"damageInflict"`
}

type SpellTime struct {
	Number int    `json:"number"`
	Unit   string `json:"unit"`
}

type SpellRange struct {
	Type     string        `json:"type"`
	Distance SpellDistance `json:"disance"`
}

type SpellDistance struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

type SpellClasses struct {
	ClassList []Class `json:"fromClassList"`
}

type Class struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}
