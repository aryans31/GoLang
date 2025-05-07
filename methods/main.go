package main

import (
	"fmt"
)

func main(){

	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: ""}
	
	fmt.Println(person.Greet())
	fmt.Println("Before update:", person)
	person.UpdateAge(31)
	fmt.Println("After update:", person)
}

type Person struct {
	Name string
	Age  int
	Email string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}