package main

import (
	"fmt"
	"time"
)

var balance int

func main() {

	for {
		fmt.Println("#################################")
		fmt.Println("Welcome to Arasaka banking system")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("#################################")
		fmt.Println("Please select what you want to do: ")

		userInput()

		time.Sleep(3 * time.Second)

	}

}

func userInput() {

	var userInput uint

	fmt.Scan(&userInput)

	switch userInput {
	case 1:
		fmt.Printf("Your balance is: %v \n", balance)
	case 2:
		var amount int
		fmt.Println("How much do you want to deposit: ")
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Println("You can only deposit more than 0")
		} else {
			fmt.Printf("You deposited: %v \n", amount)
			balance += amount
			fmt.Printf("Your new balance is: %v \n", balance)
		}
	case 3:
		var amount int
		fmt.Println("How much do you want to withdraw: ")
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Println("You can only withdraw more than 0")
		} else {
			fmt.Printf("You withdrew: %v \n", amount)
			balance -= amount
			fmt.Printf("Your new balance is: %v \n", balance)
		}
	default:
		fmt.Println("Please select 1,2 or 3")
	}
}
// done!