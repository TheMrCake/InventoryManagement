package main

import (
	"fmt"
	"strings"
)

func Add() {
	fmt.Printf("Please enter a barcode to add one to the quantity of that item. Type exit to return to menu.")
	fmt.Println()
	var input string
	fmt.Scanln(&input)
	for strings.ToLower(input) != "exit" {
		if it, ok := Inventory[input]; ok {
			arith(it, 1)
		} else {
			fmt.Println("fuck you you stupid bitch") // TODO: lmao change this
		}
		fmt.Scanln(&input)
	}
}

func Remove() {
	fmt.Printf("Please enter a barcode to remove one to the quantity of that item. Type exit to return to menu.")
	fmt.Println()
	var input string
	fmt.Scanln(&input)
	for strings.ToLower(input) != "exit" {
		if it, ok := Inventory[input]; ok {
			arith(it, -1)
		} else {
			fmt.Println("fuck you you stupid bitch") // TODO: lmao change this
		}
		fmt.Scanln(&input)
	}

	// TODO: Using MinQuantity to warn.
	// TODO: Negative quantity error.
}

func arith(i *item, n int) {
	i.Quantity += n
	fmt.Printf("Added %v to %s", n, i.Name)
	fmt.Println()
}

func Help() {
	fmt.Printf("Enter a barcode or scan an item to display its properties, or enter a command.\n - add: \n - remove: \n - new: \n - import: \n - export: \n - list: \n - orderlist: \n - help: \n")
}
