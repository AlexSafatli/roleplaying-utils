package main

import "os"

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	var vaults []*Vault
	for _, path := range os.Args[1:] {
		vault, err := ReadVault(path)
		if err != nil {
			panic(err)
		}
		vaults = append(vaults, vault)
	}
	for _, vault := range vaults {
		err := SummarizeVault(vault)
		if err != nil {
			panic(err)
		}
	}
}
