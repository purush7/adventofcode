package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("It should have only 2 arguments including main.go")
	}
	inputFile := os.Args[1]
	var ans = int64(0)
	inputBytes, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inputBytes)

	var inputString string
	var inputNumber int
	var prevNumber int = -1
	for scanner.Scan() {
		inputString = scanner.Text()
		inputNumber, err = strconv.Atoi(inputString)
		if err != nil {
			log.Fatal(err)
		}
		if prevNumber == -1 {
			prevNumber = inputNumber
			continue
		}
		if inputNumber > prevNumber {
			ans++
		}
		prevNumber = inputNumber
	}
	fmt.Println(ans)

}
