package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func IncreasingCount(scanner *bufio.Scanner) int64 {
	var ans = int64(0)
	var inputString string
	var inputNumber int
	var err error
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

	return ans
}

func IncreasingCountInWindow(scanner *bufio.Scanner) int64 {
	var ans = int64(0)
	var inputString string
	inputNumbers := []int{}
	var inputNumber int
	var err error

	for scanner.Scan() {
		inputString = scanner.Text()
		inputNumber, err = strconv.Atoi(inputString)
		if err != nil {
			log.Fatal(err)
		}
		inputNumbers = append(inputNumbers, inputNumber)
	}

	if len(inputNumbers) <= 3 {
		log.Fatal("Given only 3 numbers so we can't compare the windows as only 1 is formed")
	}

	for i := 3; i < len(inputNumbers); i++ {
		if inputNumbers[i] > inputNumbers[i-3] {
			ans++
		}
	}

	return ans
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("It should have only 2 arguments including main.go")
	}
	inputFile := os.Args[1]
	inputBytes, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inputBytes)
	fmt.Println(IncreasingCount(scanner))
	inputBytes, err = os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(inputBytes)
	fmt.Println(IncreasingCountInWindow(scanner))
}
