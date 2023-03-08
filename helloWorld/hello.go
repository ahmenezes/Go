package main

// part 1 - import fmt for basic operations
// part 2 - import a package from https://pkg.go.dev/rsc.io/quote@v1.5.2

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// part 1 - first hello world
	fmt.Println("Hello, World!")
	// part 2 - quote import usage
	fmt.Println(quote.Go())
}
