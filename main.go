package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("called 'fail' with no arguments")
		fmt.Println("Usage: fail <message>")
		fmt.Println("Example: fail 'something went wrong'")
		fmt.Println("Exiting with code 1")
		os.Exit(1)
	}
	message := os.Args[1]
	fmt.Println(message)
	os.Exit(1)
}
