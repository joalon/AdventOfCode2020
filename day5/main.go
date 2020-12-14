package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	SolvePart1()
	SolvePart2()
}

func SolvePart1() {

	input := ReadAndParse("input.txt")

	highestSeatId := math.MinInt64
	for i := 0; i < len(input); i++ {

		row := GetRow(input[i][0:7])
		column := GetColumn(input[i][7:10])

		seatId := row*8 + column
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	fmt.Printf("Highest seat ID is %d\n", highestSeatId)
}

func GetRow(seat string) int {

	left := 0
	right := 128
	mid := left + (right-left)/2

	for i := 0; i < len(seat); i++ {
		switch string(seat[i]) {
		case "F":
			right = mid
		case "B":
			left = mid
		}
		mid = left + (right-left)/2
	}

	return left
}

func GetColumn(seat string) int {
	left := 0
	right := 7
	mid := left + (right-left)/2

	for i := 0; i < len(seat); i++ {
		switch string(seat[i]) {
		case "L":
			right = mid
		case "R":
			left = mid
		}
		mid = left + (right-left)/2
	}

	return right
}

func SolvePart2() {
	input := ReadAndParse("input.txt")
	ids := []int{}

	for _, seat := range input {
		id := GetId(GetRow(seat[0:7]), GetColumn(seat[7:10]))
		ids = append(ids, id)
	}

	sort.Ints(ids)

	for i := 1; i < len(ids)-1; i++ {
		if ids[i-1] != ids[i]-1 || ids[i+1] != ids[i]+1 {
			fmt.Println("Found!", ids[i]+1)
			break
		}
	}
}

func GetId(row int, column int) int {
	return row*8 + column
}

func ReadAndParse(file string) []string {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	result := make([]string, 0)

	for s.Scan() {
		result = append(result, s.Text())
	}

	return result
}
