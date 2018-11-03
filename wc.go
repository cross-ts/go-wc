package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	flag.Parse()

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
