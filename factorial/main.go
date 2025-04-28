package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello World from Gohuda")


	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("You entered:", reader)
	reader1, _ := strconv.Atoi((reader))
	fmt.Printf("you entered:%T", reader1)
}
