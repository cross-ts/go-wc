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
			file, err := os.Open(fileName)
			defer file.Close()
			if err != nil {
				fmt.Printf("wc: %v: unknown file\n", fileName)
				return
			}
			line, word, size, name := wc(file)
			fmt.Printf("%v %v %v %v\n", line, word, size, name)
		}(fileName)
	}
	wg.Wait()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func wc(file *os.File) (line int, word int, size int64, name string) {
	stat, err := file.Stat()
	check(err)
	size, name = stat.Size(), stat.Name()

	scanner := bufio.NewScanner(file)
	line, word = 0, 0
	for scanner.Scan() {
		line++
		word += len(strings.Fields(scanner.Text()))
	}
	return
}
