package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strconv"
    "sort"
)

func main() {
	SolvePart1()
	SolvePart2()
}

func SolvePart1() {
    input := ReadAndParse("input.txt")
    sort.Ints(input)

    input = append(input, input[len(input)-1]+3)

    oneJoltDiff := 1
    threeJoltDiff :=0

    for i := 1; i < len(input); i++ {
        switch input[i] - input[i-1] {
        case 1:
            oneJoltDiff++
        case 3:
            threeJoltDiff++
        }

    }

    fmt.Println("Part 1:", oneJoltDiff * threeJoltDiff)
}

func SolvePart2() {
}


func ReadAndParse(file string) []int {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	result := make([]int, 0)

	for s.Scan() {
        parsedInt, _ := strconv.Atoi(s.Text())
        result = append(result, parsedInt)
	}

	return result
}
