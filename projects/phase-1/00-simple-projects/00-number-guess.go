package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	fmt.Println("Welcome to the Number Guessing Game")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	for {
		fmt.Print("Enter your guess ")
		fmt.Scanln(&guess)
		attempts++

		if guess < target {
			fmt.Println("low Try again.")
		} else if guess > target {
			fmt.Println("high Try again.")
		} else {
			fmt.Println("You guessed it!")
			fmt.Println("Number was:", target)
			fmt.Println("Attempts:", attempts)
			break
		}
	}
}
