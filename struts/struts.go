package main

import "fmt"

func main(){

	aryan :=User{"Aryan", 20, ""}
	fmt.Println(aryan)

}

type User struct {
	Name string
	Age  int
	Email string
	
}