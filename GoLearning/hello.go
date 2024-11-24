package main

import (
	"fmt"

	"rsc.io/quote"
)

// I tried Version 3 but it didn't work because when i ran
// go tidy command it download the latest version by default.
func main() {
	fmt.Println(quote.Hello())
}
