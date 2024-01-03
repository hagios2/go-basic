package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mybill := createBil()

	promptOptions(mybill)
	fmt.Println(mybill)
}

func createBil() bill {

	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
} 

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')
	return strings.TrimSpace(input), error
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err  := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		// b.updateTip(&p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err  := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		// b.updateTip(&p)
		fmt.Println("tip added - ", tip)

		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)

	default: 
		fmt.Println("That was not a valid option")
		promptOptions(b) 
	}
}


