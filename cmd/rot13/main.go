// Package main provides a rot13 executable.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mcandre/go-rotate"
)

// main is the entrypoint for launching this application.
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(rotate.Rot13(scanner.Text()))
	}
}
