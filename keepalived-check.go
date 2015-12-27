package main

import (
	"./keepalived"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		scanner := new(keepalived.Scanner)
		scanner.Init(string(keepalived.Load(arg)))
		for _, statement := range keepalived.Parse(scanner) {
			fmt.Println(statement)
			// fmt.Printf("%+v\n", statement)
		}
	}
}
