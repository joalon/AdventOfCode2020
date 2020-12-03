
package main

import (
    "os"
    "bufio"
    "log"
    "fmt"
)

func main() {
    SolvePart1()
    SolvePart2()
}

func SolvePart1() {
    input := ReadAndParse("input.txt")

    x := 0
    y := 0

    treesHit := 0

    for y < len(input) {
        if string(input[y][x]) == "#" {
            treesHit++
        }

        y += 1
        x += 3
        x = x % len(input[0])
    }

    fmt.Println("Hitting", treesHit, "trees in part 1")
}

type Tuple struct {
    X_speed int
    Y_speed int
}

func SolvePart2() {
    input := ReadAndParse("input.txt")

    speed_tuples := []Tuple{
        Tuple{
            X_speed: 1,
            Y_speed: 1,
        },
        Tuple{
            X_speed: 3,
            Y_speed: 1,
        },
        Tuple{
            X_speed: 5,
            Y_speed: 1,
        },
        Tuple{
            X_speed: 7,
            Y_speed: 1,
        },
        Tuple{
            X_speed: 1,
            Y_speed: 2,
        },
    }

    result := make([]int, 0)

    for _, tuple := range speed_tuples {

        x := 0
        y := 0

        treesHit := 0

        for y < len(input) {
            if string(input[y][x]) == "#" {
                treesHit++
            }

            y += tuple.Y_speed
            x += tuple.X_speed
            x = x % len(input[0])
        }

        result = append(result, treesHit)
    }

    product := 1
    for _, i := range result {
        product *= i
    }

    fmt.Printf("Product is %d in part 2\n", product)
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
