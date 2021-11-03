package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - " + b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save the bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added -", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		fmt.Println("Tip added -", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Save the bill -", b.name)
	default:
		fmt.Println("Invalid option!")
		promptOptions(b)
	}
}

func main() {
	//mybill := newBill("mario's bill")
	//
	//fmt.Println(mybill.format())
	//
	//mybill.addItem("onion soup", 4.5)
	//mybill.addItem("veg pie", 8.95)
	//mybill.addItem("toffee pudding", 4.95)
	//mybill.addItem("coffee", 3.25)
	//mybill.updateTip(10)
	//
	//fmt.Println(mybill.format())

	mybill := createBill()
	promptOptions(mybill)
}
