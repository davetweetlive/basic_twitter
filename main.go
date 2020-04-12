package main

import (
	"fmt"
	"os"
)

var (
	name   string
	course string
	module float32
)

func main() {
	// To get Os username
	name = os.Getenv("USERNAME")
	fmt.Println(name)

	// os environment variable
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
