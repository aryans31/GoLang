package main

import "fmt"


func callbackFunction() {
	fmt.Println("This is a callback function")
}
func higherOrderFunction(callback func()) {
	fmt.Println("This is a higher-order function")
	callback()
}

func main(){

	fmt.Println("hello form main function")

	higherOrderFunction(callbackFunction)
}