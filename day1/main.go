package main

import (
    "os"
    "bufio"
    "log"
    "fmt"
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

    right := len(input) - 1
    left := 0

    result := input[right] + input[left]

    for result != 2020 {
        if result > 2020 {
            right--
        } else {
            left++
        }
        result = input[right] + input[left]
    }

    fmt.Println(input[right], " * ", input[left], " = ", input[right] * input[left])
}

func SolvePart2() {
    input := ReadAndParse("input.txt")

    // I spent waaay to much time trying to be clever. Let's bruteforce it instead:
    for _, x := range input {
        for _, y := range input {
            for _, z := range input {
                if x + y + z == 2020 {
                    fmt.Println(x, " * " , y, " * ", z, " = ", x * y * z)
                    goto done
                }
            }
        }
    }
done:
    return
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
        i, _ := strconv.Atoi(s.Text())
        result = append(result, i)
    }

    return result
}
