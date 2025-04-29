
package main
import (
	"fmt"
)
func main()  {


	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		if k >= 10 {
			break
		}
		fmt.Println(k)
		k++
	}
	
	// for range loop
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// nested for loop
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i:", i, "j:", j)
		}
	}
	// break and continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // skip the rest of the loop when i is 5
		}
		if i == 8 {
			break // exit the loop when i is 8
		}
		fmt.Println(i)
	}
	// labeled break
outerLoop:
	for i := 0; i < 3; i++ {	
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outerLoop // exit the outer loop when i is 1 and j is 1
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	
}