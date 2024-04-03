package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	if argsWithoutProg[0] == "-w" {
		readFile, err := os.ReadFile(myFile)
		if err != nil {
			fmt.Println(err)
		}
		result := strings.Fields(string(readFile))
		fmt.Println(len(result), myFile)
	}
	if argsWithoutProg[0] == "-m" {
		readFile, err := os.ReadFile(myFile)
		if err != nil {
			fmt.Println(err)
		}
		data := string(readFile)
		result :=  len(([]rune)(data))
		fmt.Println(result, myFile)
	}
}
