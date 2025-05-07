package main

import (
	"fmt"
)

func main(){

	fmt.Println("Hello from different functions")

 var a,b,c int
 fmt.Println("Enter two numbers")
fmt.Scan(&a,&b)
 c = add(a,b)
 fmt.Println("Sum is",c)
}

func add(a ...int) int { //... variadic function

	return a[0] + a[1]
}

