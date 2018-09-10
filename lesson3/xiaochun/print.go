package main

import "fmt"

func main() {
	str1 := "hello"
    doc := `
即使换行也不影响
可以验证一下
类似Python的'''
`
	fmt.Println(str1,doc)
}
