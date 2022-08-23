package main

import (
	"fmt"
)

func Add(i *item, n int) {
	i.Quantity += n
	fmt.Printf("Added %v, to %s", n, i.Name)

}

func Remove(i *item, n int) {
	i.Quantity -= n
	// TODO: Using MinQuantity to warn.
	// TODO: Negative quantity error.
}

func Help() {
	fmt.Printf("Enter a barcode or scan an item to display its properties, or enter a command.\n - add: \n - remove: \n - new: \n - import: \n - export: \n - list: \n - orderlist: \n - help: ")
}
