package main

import (
	"fmt"
)

func main() {
	s := "hello中文"
	for i, arg := range s {
		fmt.Printf("%d %c\n", i, arg)
	}

}
