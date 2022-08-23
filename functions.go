package main

import "fmt"

func add(i *item, n int) {
	i.Quantity += n
	fmt.Printf("Added %v, to %s", n, i.Name)

}

func remove(i *item, n int) {
	i.Quantity -= n
	// TODO: Using MinQuantity to warn.
}
