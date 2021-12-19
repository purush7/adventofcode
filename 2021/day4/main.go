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

func markNumber(board *[]boardType, number int, keyPositions *[]int) {
	for ind, eachBoard := range *board {
		for k, v := range *eachBoard {
			if v.value != number || v.crossed {
				continue
			}
			v.crossed = true
			(*keyPositions)[ind] = k
			break
		}
	}
}

func checkBingo(board *[]boardType, keyPositions *[]int, bingoBoards *map[int]bool, part1 bool) (int, int) {
	sm, id := -1, -1
	for ind, eachBoard := range *board {
		if (*bingoBoards)[ind] {
			continue
		}
		bingoRow, bingoCol := true, true
		k := (*keyPositions)[ind]
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
			sm = boardSum(board, ind)
			(*bingoBoards)[ind] = true
			id = ind
			if part1 {
				return sm, id
			}
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
			sm = boardSum(board, ind)
			(*bingoBoards)[ind] = true
			id = ind
			if part1 {
				return sm, id
			}
		}
	}
	return sm, id
}

func part1(numbersArray *[]int, board *[]boardType, keyPositions *[]int, bingoBoards *map[int]bool) int {
	for _, numb := range *numbersArray {
		markNumber(board, numb, keyPositions)
		ans, _ := checkBingo(board, keyPositions, bingoBoards, true)
		if ans != -1 {
			return ans * numb
		}
	}
	return -1
}

func part2(numbersArray *[]int, board *[]boardType, keyPositions *[]int, bingoBoards *map[int]bool) int {
	sol := -1
	for _, numb := range *numbersArray {
		markNumber(board, numb, keyPositions)
		ans, boardId := checkBingo(board, keyPositions, bingoBoards, false)
		if ans != -1 && boardId != -1 {
			if ans != 0 {
				sol = ans * numb
			}

		}
	}
	return sol
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

	noOfBoards := len(board)

	if noOfBoards == 0 {
		log.Fatal("no of Boards is zero")
	}

	keyPositions := make([]int, noOfBoards)

	bingoBoards := make(map[int]bool, 0)

	for boardIndex := 0; boardIndex < noOfBoards; boardIndex++ {
		bingoBoards[boardIndex] = false
	}

	fmt.Println("part1 ans: ", part1(numbersArray, &board, &keyPositions, &bingoBoards))
	refershBoard(&board)
	fmt.Println("part2 ans: ", part2(numbersArray, &board, &keyPositions, &bingoBoards))
}
