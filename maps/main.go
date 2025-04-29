package main

import "fmt"

func main(){

	//crate a map
	map1 := make(map[string]string)

	//add elements to the map
	map1["js"] = "JavaScript"
	map1["RB"] = "Ruby"
	map1["PY"] = "Python"

	fmt.Println("list of all languages", map1)

	delete(map1, "RB") //delete an element from the map
	fmt.Println("list of all languages after deleting Ruby", map1)

	//print the map
	for key,value := range map1{
		println(key,value)
	}
}