package main

type item struct {
	Name, Description, Barcode string
	Quantity                   int
	MinQuantity                int
}

func newItem(n string, des string, bd string, q int, min int) *item {
	i := item{Name: n, Description: des, Barcode: bd, Quantity: q, MinQuantity: min}
	return &i
}
