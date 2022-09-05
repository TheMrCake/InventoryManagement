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

		switch in := strings.ToLower(input); in {
		case "add":
			Add()
		case "remove":
			Remove()
		default:
			if it, ok := Inventory[in]; ok {
				fmt.Println(it)
			} else {
				fmt.Println("fuck you you stupid bitch") // TODO: lmao change this
			}
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
