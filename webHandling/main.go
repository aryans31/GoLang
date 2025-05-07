package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {

	fmt.Println("Starting web server on :8080")
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Headers:", response.Header)
	fmt.Println("Response Body:")
	fmt.Printf("%T", response)

	dataBytes,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Data in response body:", string(dataBytes))
}
	
