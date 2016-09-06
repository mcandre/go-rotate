package main

import (
	"bufio"
	"fmt"
	"github.com/mcandre/go-rotate"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(rotate.Rot13(scanner.Text()))
	}
}
