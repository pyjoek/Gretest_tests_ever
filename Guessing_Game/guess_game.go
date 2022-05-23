package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for{
		min, max := 111, 999
		guess := rand.Intn(max-min) + min
		var user int
		fmt.Print("Enter your guess between 1, 100: ")
		fmt.Scan(&user)
		if user == guess{
			fmt.Printf("You won the guess is %v\n",guess)
			break
		}else{
			fmt.Printf("You missed it, the guess is %v\n",guess)
		}
	}
}
