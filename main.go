package main

import (
	_ "embed"
	"fmt"
)

// must be create outside function
//
//go:embed message.txt
var message string

func main() {
	fmt.Println(message)
}
