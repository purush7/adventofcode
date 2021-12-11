package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Dive(scanner *bufio.Scanner) int64 {
	var inputLine string
	var val int
	var err error
	var hor, vert int64

	for scanner.Scan() {
		inputLine = scanner.Text()
		inp := strings.Split(inputLine, " ")
		val, err = strconv.Atoi(inp[1])
		if err != nil {
			log.Fatal(err)
		}
		switch inp[0] {
		case "forward":
			hor += int64(val)
		case "up":
			vert -= int64(val)
		case "down":
			vert += int64(val)
		}
	}
	if hor > 0 && vert > 0 {
		return hor * vert
	}

	log.Fatal("any one of the horizatal or the depth is negative")

	return int64(0)
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
	fmt.Println(Dive(scanner))
}
