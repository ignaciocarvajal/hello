package main

import (
	"fmt"

	"github.com/ignaciocarvajal/hello/greet"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello", greet.Italian())
	fmt.Println("Hello", quote.Hello())
}
