package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	for i, v := range os.Args {
		fmt.Printf("%d ==> %v\n", i, v)
	}
}
