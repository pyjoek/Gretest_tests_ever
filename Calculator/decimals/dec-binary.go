package main

import "fmt"

func main(){
	var i int
	for i = 1;i > 1;i++ {
		binary := []int{}
		var n = 20
		var a = 0.5
		var b = 0.9
		var p = n % 2
		if p >= a && p <= b{
			p = 1
		}else if p >= 1.0 && p <= 1.4{
			p = 1
			n = n/2
		}
	fmt.Println(binary)
	}
}
