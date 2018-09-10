package main

import "fmt"

func main() {
		s1 := "hello" + "world"
		s2 := "helloworld"

		if s1 == s2 {
				fmt.Println("equal")
		}
		
		fmt.Println(0, len(s1)-1)
		var c1 byte
		c1 = s1[0]
		fmt.Println(s1, s2, c1)
		fmt.Printf("%d %c\n", c1, c1)
		s3 := s1[1:7]
		fmt.Println(s3)

		var b byte
		for b = 0; b < 177; b++ {
				fmt.Printf("%d %c\n", b, b)
		}
		
		array := []byte(s1)
		fmt.Println(array)
		array[0] = 'H'
		s1 = string(array)		
		fmt.Println(s1)

		fmt.Println(0xA)

		fmt.Println('a' + ('H' - 'h'))
}
