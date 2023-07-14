package main

import "fmt"

func main() {
	maxNumber := 21
	for i := 1; i < maxNumber; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if (i%3 != 0) && (i%5 != 0) {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}
