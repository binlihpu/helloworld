package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	name2 := *flag.String("name2", "world", "test")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
	fmt.Printf("Hello, %v!\n", name2)
}
