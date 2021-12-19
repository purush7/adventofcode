package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type elementType struct {
	value   int
	crossed bool
}

type boardType *map[int]*elementType

func refershBoard(board *[]boardType) {
	for _, eachBoard := range *board {
		for _, val := range *eachBoard {
			val.crossed = false
		}
	}
}

func boardSum(board *[]boardType, index int) int {
	sum := 0
	for _, v := range *(*board)[index] {
		if !v.crossed {
			sum += v.value
		}
	}
	return sum
}

func checkBingo(board *[]boardType, number int) int {
	for ind, eachBoard := range *board {
		bingoRow, bingoCol := true, true
		for k, v := range *eachBoard {
			if v.value != number || v.crossed {
				continue
			}
			v.crossed = true

			//check for bingo

			// rows
			r := k / 5
			for i := 0; i < 5; i++ {
				if (*eachBoard)[5*r+i].crossed {
					continue
				}
				bingoRow = false
				break
			}

			if bingoRow {
				sum := boardSum(board, ind)
				return sum * number
			}

			// cols
			c := k % 5
			for i := 0; i < 5; i++ {
				if (*eachBoard)[i*5+c].crossed {
					continue
				}
				bingoCol = false
				break
			}

			if bingoCol {
				sum := boardSum(board, ind)
				return sum * number
			}
		}
	}

	return -1
}

func part1(numbersArray *[]int, board *[]boardType) int {
	for _, numb := range *numbersArray {
		ans := checkBingo(board, numb)
		if ans != -1 {
			return ans
		}
	}
	return -1
}

func main() {

	var inputLine string
	var numbersArray = new([]int)

	if len(os.Args) != 2 {
		log.Fatal("It should have only 2 arguments including main.go")
	}
	inputFile := os.Args[1]
	inputBytes, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inputBytes)

	// number array
	scanner.Scan()
	inputLine = scanner.Text()
	stringArray := strings.Split(inputLine, ",")
	for _, val := range stringArray {
		number, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		*numbersArray = append(*numbersArray, number)
	}

	board := make([]boardType, 0)

	for scanner.Scan() {
		// 5x5
		ele := make(map[int]*elementType)
		for x := 0; x < 5; x++ {
			scanner.Scan()
			inputLine = scanner.Text()
			stringArray := strings.Split(inputLine, " ")
			y := 0
			for _, val := range stringArray {
				if val == "" {
					continue
				}
				nm, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal(err)
				}
				elementNode := new(elementType)
				elementNode.value = nm
				elementNode.crossed = false
				ele[x*5+y] = elementNode
				y++
			}
		}
		board = append(board, &ele)
	}
	fmt.Println("part1 ans: ", part1(numbersArray, &board))
}
