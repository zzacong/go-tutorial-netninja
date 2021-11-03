package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := fmt.Sprintf("%v bill breakdown: \n", b.name)
	total := .0

	for k, v := range b.items {
		// %-25 : pad with spaces on the right rather than the left (left-justify the field)
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%.2f \n", "total:", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/" + b.name + "-bill.txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
