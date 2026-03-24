package main

import (
	"fmt"
)

// Write a Go program that does the following:

// 1. Declares a struct called Person with fields: Name (string), Age (int)
// 2. Creates two Person values
// 3. Uses a for loop to print each person's name and age
// 4. Uses an if statement to print "Senior" if age > 40, else "Junior"
// 5. Uses a map to count how many times each label (Senior/Junior) appears
// 6. Prints the final count

// No functions other than main. No imports other than "fmt".

type person struct {
	Name string
	Age  int
}

func main() {
	person1 := person{
		Name: "Pradip",
		Age:  25,
	}

	person2 := person{
		Name: "jon",
		Age:  45,
	}

	person3 := person{
		Name: "ram",
		Age:  45,
	}
	people := []person{person1, person2, person3}
	counts := make(map[string]int)
	for _, p := range people {
		if p.Age > 40 {
			fmt.Println("Senior")
			counts["Senior"]++
		} else {
			fmt.Println("Junior")
			counts["Junior"]++
		}
	}
	fmt.Println("Final count")
	for label, count := range counts {
		fmt.Println(label+":", count)
	}
}
