package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Options struct {
	c bool
	m bool
	l bool
}

var options = Options{c: false}

func parse() {
	flag.BoolVar(&options.c, "c", false, "print the byte counts")
	flag.BoolVar(&options.m, "m", false, "print the character counts")
	flag.BoolVar(&options.l, "l", false, "print the newline counts")
	flag.Parse()
}

func main() {
	parse()

	wg := sync.WaitGroup{}
	for _, fileName := range flag.Args() {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			wc(fileName)
		}(fileName)
	}
	wg.Wait()
}

func wc(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("wc: %v: unknown file\n", fileName)
		return
	}

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size, name := stat.Size(), stat.Name()

	scanner := bufio.NewScanner(file)
	line, word := 0, 0
	for scanner.Scan() {
		line++
		word += len(strings.Fields(scanner.Text()))
	}
	fmt.Printf("%v %v %v %v\n", line, word, size, name)
	return
}
