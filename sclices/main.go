package main

import "fmt"

func main(){
fmt.Println("hello sclices")

sclice1 := []int{1,2,3,4,5}

fmt.Println(sclice1)
fmt.Println("Length of sclice1:", len(sclice1))
fmt.Println("Capacity of sclice1:", cap(sclice1))
fmt.Println("First element of sclice1:", sclice1[0:3])
sclice1 = append(sclice1, 6)
fmt.Println("After appending 6:", sclice1)
fmt.Println("deleting 2nd element")
sclice1 = append(sclice1[:1], sclice1[2:]...)
}