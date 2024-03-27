package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	myFile := os.Args[2]
    argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "-c" {
		MyFile, err := os.Stat(myFile)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(MyFile.Size(), myFile)
	}
	if argsWithoutProg[0] == "-l" {
		readFile, err := os.Open(myFile)
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(readFile)
		var count int
		for scanner.Scan() {
		   count++
		}
		fmt.Println(count)
		defer readFile.Close()
	}
}
