package main

type item struct {
	Name, Description, Barcode string
	Quantity                   int16
	MinQuantity                int16
}

func newItem(n string, des string, bd string, q int16, min int16) *item {
	i := item{Name: n, Description: des, Barcode: bd, Quantity: q, MinQuantity: min}
	return &i
}

func add()
