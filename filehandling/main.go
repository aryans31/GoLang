package main

import (
	"fmt"
	"io"
	"os"
)

	func main(){

		message := "This is a new file"

		file,err :=os.Create("./test.txt")

		if err != nil {
			fmt.Println("Error creating file:", err)
			panic(err)

		}

		length ,err := io.WriteString(file, message)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			panic(err)
		}
		fmt.Println("Length of file:", length)
		fmt.Println("File created successfully")

		defer file.Close()
		readFile("./test.txt")

	}

	func readFile(fileName string){
		dataByte, err:= os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			panic(err)
		}
		fmt.Println("Data in file:", string(dataByte))
	}