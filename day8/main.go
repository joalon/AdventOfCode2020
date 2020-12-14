package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "strings"
    "strconv"
)

func main() {
	SolvePart1()
	SolvePart2()
}

func SolvePart1() {
	input := ReadAndParse("input.txt")

    accHistory := make(map[int]int)
    i, acc := 0, 0

    for {
        if _, ok := accHistory[i]; ok {
            break
        }
        accHistory[i] = acc

        switch input[i].op {
        case "jmp":
            i += input[i].arg
        case "acc":
            acc += input[i].arg
            i++
        case "nop":
            i++
        }
    }

    fmt.Println("Part 1:", acc)
}

func SolvePart2() {
//	input := ReadAndParse("input.txt")
//
//    acc := 0
//    i := 0
//    for {
//        switch input[i].op {
//        case "jmp":
//            i += input[i].arg
//        case "acc":
//            acc += input[i].arg
//            i++
//        case "nop":
//            i++
//        }
//    }
//
//    fmt.Println("Part 2:", acc)
}

type Instruction struct {
    op  string
    arg int
}


func ReadAndParse(file string) []Instruction {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	result := make([]Instruction, 0)

	for s.Scan() {
        line := strings.Split(s.Text(), " ")
        parsedArg, _ := strconv.Atoi(line[1])
        inst := Instruction{op: line[0], arg: parsedArg}
        result = append(result, inst)
	}

	return result
}
