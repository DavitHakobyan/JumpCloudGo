package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please execute the myLib_test.go")
	go Davit()
}

func Davit() {
	for {
		fmt.Println("Hello")
	}

}