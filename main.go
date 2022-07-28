package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args
	res := 0
	if len(args) > 1 {
		str := args[1]
		if len(str) > 0 {
			res = strings.Count(str, " ") + 1
		}
	}
	fmt.Println(res)
}
