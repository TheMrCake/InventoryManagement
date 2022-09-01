package main

import (
	"fmt"
	"strings"
)

var Inventory = make(map[string]*item)

func main() {
	Init()
	for {
		Help()

		var input string
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "add":
			Add()

		}
	}

	// TODO: Export ExportAutoSave
}

func Init() {
	ImportAutoSave()

	// TODO: Run Unit tests (lmao like I'm ever making those)

}

func ImportAutoSave() {
	ItemOne := NewItem("Item1", "I am an Item", 3, 2)
	Inventory["12345-678-901"] = ItemOne
	// TODO: Implement this function

}
