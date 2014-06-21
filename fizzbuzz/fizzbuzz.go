package main

import "fmt"
import "math"

func main() {
	i := 1.0
	for i < 101.0 {
		if math.Mod(i, 15) == 0 {
			fmt.Println("Fizzbuzz")
		} else if math.Mod(i, 3) == 0 {
			fmt.Println("Fizz")
		} else if math.Mod(i, 5) == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}
