package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	SolvePart1()
	SolvePart2()
}

func SolvePart1() {
	input := ReadAndParse("input.txt")

    numberOfAnswers:= 0

    for _, declaration := range input {
        numberOfAnswers += len(declaration.answers)
    }

    fmt.Println("Part 1:", numberOfAnswers)
}

func SolvePart2() {
	input := ReadAndParse("input.txt")

    numberOfAnswers := 0

    for _, declaration := range input {
        for _, v := range declaration.answers {
            if v == declaration.numberOfPeople {
                numberOfAnswers++
            }
        }
    }

    fmt.Println("Part 2:", numberOfAnswers)
}

type CustomsDeclaration struct {
    numberOfPeople  int
    answers         map[string]int
}

func ReadAndParse(file string) []CustomsDeclaration {
	fh, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	s := bufio.NewScanner(fh)
	result := make([]CustomsDeclaration, 0)

    current := CustomsDeclaration{numberOfPeople: 0, answers: make(map[string]int)}

	for s.Scan() {
        line := s.Text()
        if line == "" {
            result = append(result, current)
            current = CustomsDeclaration{numberOfPeople: 0, answers: make(map[string]int)}
            continue
        }

        current.numberOfPeople++

        for _, char := range line {
            current.answers[string(char)] += 1
        }
	}

	return result
}
