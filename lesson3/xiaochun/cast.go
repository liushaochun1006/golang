package main

import (
		"fmt"
		"strconv"
		"math/rand"
		"time"
)

func main() {
		var n int
		var f float32
		n = 10
		f = float32(n) / 3
		fmt.Println(f * 3)
		n = int(f * 3)

		fmt.Println(f, n)

		var n1 int64
		n1 = 1024004
		var n2 int8
		n2 = int8(n1)
		fmt.Println(n1, n2)

		n1 = 1024128
		n2 = int8(n1)
		fmt.Println(n1, n2)

		var s string
		s = strconv.Itoa(n)
		fmt.Println(s)

		n, err := strconv.Atoi("abc123")
		if err != nil {
				fmt.Println("error",err)
		}
		fmt.Println(n)

		var x int64
		rand.Seed(time.Now().Unix())
		x = rand.Int63()
		s = strconv.FormatInt(x, 2)
		fmt.Println(s)
}
