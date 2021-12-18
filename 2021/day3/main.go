package main

import (
	"bufio"
	"container/list"
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

func BinaryPart2(scanner *bufio.Scanner, sz int) int64 {
	var inputLine string
	var ind = int64(0)
	var err error
	// var hor, vert int64
	var mp = make(map[int64]string)
	var strlen int

	for scanner.Scan() {
		inputLine = scanner.Text()
		mp[ind] = inputLine
		strlen = len(inputLine)
		ind++
	}

	var delKey = func(ls *list.List, mp *map[int64]string) {
		for ls.Len() > 0 {
			fr := ls.Front()
			ls.Remove(fr)
			delete(*mp, fr.Value.(int64))
		}
	}

	var searchBit = func(mp map[int64]string, maxMin int64, strlen int) string {
		var ans string

		if maxMin == 1 {
			for i := 0; i < strlen && len(mp) > 0; i++ {
				oneList := list.New()
				zeroList := list.New()
				for k, v := range mp {
					if v[i] == '1' {
						oneList.PushBack(k)
					} else {
						zeroList.PushBack(k)
					}
				}
				if oneList.Len() < zeroList.Len() {
					ans = ans + "0"
					delKey(oneList, &mp)
				} else {
					ans = ans + "1"
					delKey(zeroList, &mp)
				}
			}
		}
		if maxMin == 0 {
			for i := 0; i < strlen && len(mp) > 0; i++ {
				oneList := list.New()
				zeroList := list.New()
				for k, v := range mp {
					if v[i] == '1' {
						oneList.PushBack(k)
					} else {
						zeroList.PushBack(k)
					}
				}
				if len(mp) == 1 {
					for kk := range mp {
						ans = mp[kk]
					}
					break
				}
				if zeroList.Len() <= oneList.Len() {
					ans = ans + "0"
					delKey(oneList, &mp)
				} else {
					ans = ans + "1"
					delKey(zeroList, &mp)
				}
			}
		}

		return ans
	}

	mp2 := make(map[int64]string, len(mp))
	for k, v := range mp {
		mp2[k] = v
	}
	o2 := searchBit(mp, 1, strlen)
	co2 := searchBit(mp2, 0, strlen)

	o2No, err := strconv.ParseInt(o2, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	co2No, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return o2No * co2No

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
	inputBytes, err = os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(inputBytes)
	fmt.Println(BinaryPart2(scanner, val))
}
