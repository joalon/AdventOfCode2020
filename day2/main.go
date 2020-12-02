
package main

import (
    "os"
    "bufio"
    "log"
    "fmt"
    "strings"
    "strconv"
)

type Password struct {
    MinRepeats int
    MaxRepeats int

    Letter string

    Password string
}

func (p Password) Valid1() bool {
    occurrenceMap := make(map[string]int)

    for _, letter := range p.Password {
        occurrenceMap[string(letter)] += 1
    }

    isValid := occurrenceMap[p.Letter] >= p.MinRepeats && occurrenceMap[p.Letter] <= p.MaxRepeats
    return isValid
}

func (p Password) Valid2() bool {
    if string(p.Password[p.MinRepeats - 1]) == p.Letter && string(p.Password[p.MaxRepeats - 1]) != p.Letter {
        return true
    } else if string(p.Password[p.MinRepeats - 1]) != p.Letter && string(p.Password[p.MaxRepeats - 1]) == p.Letter {
        return true
    }

    return false
}

func main() {
    SolvePart1()
    SolvePart2()
}

func SolvePart1() {
    input := ReadAndParse("input.txt")

    validPasswords := 0
    for _, password := range input {
        if password.Valid1() {
            validPasswords++
        }
    }

    fmt.Printf("Valid passwords: %d\n", validPasswords)
}

func SolvePart2() {
    input := ReadAndParse("input.txt")

    validPasswords := 0
    for _, password := range input {
        if password.Valid2() {
            validPasswords++
        }
    }

    fmt.Printf("Valid passwords: %d\n", validPasswords)
}

func ReadAndParse(file string) []Password {
    fh, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer fh.Close()

    s := bufio.NewScanner(fh)
    result := make([]Password, 0)

    for s.Scan() {
        split := strings.Split(s.Text(), " ")
        repeats := strings.Split(split[0], "-")
        minrepeats, _ := strconv.Atoi(repeats[0])
        maxrepeats, _ := strconv.Atoi(repeats[1])

        p := Password{
            MinRepeats: minrepeats,
            MaxRepeats: maxrepeats,
            Letter: string(split[1][0]),
            Password: split[2],
        }

        result = append(result, p)
    }

    return result
}
