/*
Write a Go program that does the following:

1. Write a function called generate(nums ...int) <-chan int
   - It launches a goroutine that sends each number into a channel
   - It closes the channel when done
   - It returns the receive-only channel

2. Write a function called square(in <-chan int) <-chan int
   - It launches a goroutine that reads from in, squares each number,
     and sends the result into a new channel
   - It closes the output channel when the input channel is closed
   - It returns the receive-only output channel

3. In main():
   - Use generate() to create a channel of numbers 1 through 5
   - Pass that channel into square()
   - Range over the result and print each squared value
   - Expected output: 1, 4, 9, 16, 25 (one per line)

4. Chain a second square() call so numbers go through squaring twice:
   generate → square → square → print
   Expected output: 1, 16, 81, 256, 625

Allowed imports: "fmt"
*/

package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	ch1 := generate(1, 2, 3, 4, 5)
	fmt.Println("Single square:")
	for v := range square(ch1) {
		fmt.Println(v)
	}

	fmt.Println("\nDouble square:")
	ch2 := generate(1, 2, 3, 4, 5)
	for v := range square(square(ch2)) {
		fmt.Println(v)
	}

}
