package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	files := flag.Args()

	fmt.Println(files)
}
