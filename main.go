package main

import (
	"fmt"
	"strings"
)

func main() {
	Init()
	for true {
		Help()

		var input string
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "add":
			Add()

		}
	}
}

func Init() {
	// TODO: Implement
}
