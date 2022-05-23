package main
import "fmt"
func add(){
	var a int
	var b int
	fmt.Println("Enter the second number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println(a + b)
}
func substract(){
	var a int
	var b int
	fmt.Println("Enter the second number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println(a - b)
}
func multiply(){
	var a int
	var b int
	fmt.Println("Enter the second number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println(a * b)
}
func divide(){
	var a int
	var b int
	fmt.Println("Enter the second number: ")
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scan(&b)
	fmt.Println(a / b)
}
func loop(){
	var do string
	fmt.Println("Do you want to do it again [ yes | y ]: ")
	fmt.Scan(&do)
	if do == "yes" || do == "y" {
		main()
	}else{
		fmt.Print("ok Good day !!!!!!")
	}
}
func main(){
	var options string
	fmt.Println("Enter the operation to be used in (add, sub, multy, divide): ")
	fmt.Scan(&options)
	if options == "add"{
		add()
		loop()
	}else if options == "sub" {
		substract()
		loop()
	}else if options == "multy" {
		multiply()
		loop()
	}else if options == "divide" {
		divide()
		loop()
	}else{
		fmt.Println("The option you choosed is not in the list Sorry !!!!!!")
		loop()
	}
}