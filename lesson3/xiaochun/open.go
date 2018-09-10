package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
		f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
				log.Fatal(err)
		}
	
		r := bufio.NewReader(f)
		for {
				line, err := r.ReadString('\n')
				if err == io.EOF {
						break
				}
				fmt.Print(line)
		
		}
		f.Close()
}
