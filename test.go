package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile(".env")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%q\n", data)

	data2, err := os.ReadFile(".env.example")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("%q\n", data2)
}
