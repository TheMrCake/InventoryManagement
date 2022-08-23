package main

import "fmt"

type item struct {
	Name, Description string
	Quantity          int
	MinQuantity       int
}

func NewItem(n string, des string, q int, min int) *item {
	i := item{Name: n, Description: des, Quantity: q, MinQuantity: min}
	return &i
}

func (i item) String() string {
	return fmt.Sprintf("Name: %s \nDescription: %s \nQuantity: %v \nMin Quantity: %v", i.Name, i.Description, i.Quantity, i.MinQuantity)
}
