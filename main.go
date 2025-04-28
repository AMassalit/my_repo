package main

import (
	"awesomeProject2/service"
	"fmt"
)

func main() {
	fmt.Printf("Hello and welcome, %s!\n", service.ServiceFunc())

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
