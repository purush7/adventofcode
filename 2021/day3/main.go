package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func BinaryPart1(scanner *bufio.Scanner, sz int) int64 {
	var inputLine string
	var len int64
	var err error
	// var hor, vert int64
	count := make([]int64, sz)
	var gamma, epsilon string

	for scanner.Scan() {
		inputLine = scanner.Text()
		for in, bin := range inputLine {
			switch bin {
			case '0':
				count[in]--
			case '1':
				count[in]++
			default:
				log.Fatal("received other than 1 and 0s")
			}
		}
		len++
	}

	for _, ele := range count {
		if ele >= 0 {
			gamma += "1"
		} else {
			gamma += "0"
		}

		if ele == len {
			epsilon += "1"
		} else if ele == -1*len {
			epsilon += "0"
		} else if ele >= 0 {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	gammaNo, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	epsilonNo, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return gammaNo * epsilonNo
}

func main() {

	if len(os.Args) != 3 {
		log.Fatal("It should have only 3 arguments including main.go")
	}
	inputFile := os.Args[1]
	val, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	inputBytes, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inputBytes)
	fmt.Println(BinaryPart1(scanner, val))
}
