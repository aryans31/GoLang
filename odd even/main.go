package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Hello World from Go")
	fmt.Println("We are doing odd even numbers")

	reader := bufio.NewReader(os.Stdin)

	num, _ := reader.ReadString('\n')
	num1, _ := strconv.Atoi(strings.TrimSpace(num))
	if num1%2 == 0 {
		fmt.Println("Number \v is even", num)
	} else {
		fmt.Println("Number \v is odd", num)
	}

}
