package main

import "fmt"

func main(){
	// for i := 0; i < 5;i++{
	// 	fmt.Printf("5 x %v = %v\n",i,5*i)
	// }
	for {
		var name string
		var age int
		var passwr string
		var gender string
		fmt.Print("Enter your name: ")
		fmt.Scan(&name)
		fmt.Print("Enter your age: ")
		fmt.Scan(&age)
		fmt.Print("Enter your gender: ")
		fmt.Scan(&gender)
		fmt.Print("Enter your password: ")
		fmt.Scan(&passwr)
		var confr string
		fmt.Print("Confirm your password: ")
		fmt.Scan(&confr)
		if (passwr == confr){
			fmt.Println("Account created successful")
			println("\n\n\n")
			fmt.Printf("#####   ACCOUNT DETAILS ########\nNAME : %v\nAGE : %v\nGENDER : %v\n",name,age,gender)
			break
		}else{
			fmt.Println("Password mismatch")
			fmt.Println("Please try again")
		}
	}
}
