package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "golang.org/x/exp/maps"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    numbers := map[string]int {
        "0": 0,
        "1": 1,
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "zero": 0,
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    numbergroup := strings.Join( maps.Keys(numbers), "|" )

    regexfirst := regexp.MustCompile("^.*?("+numbergroup+")")
    regexlast := regexp.MustCompile("^.*("+numbergroup+")")
    total := 0

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        line := scanner.Text()
        // fmt.Println(scanner.Text())
        first := numbers[ regexfirst.FindStringSubmatch(line)[1] ]
        last  := numbers[ regexlast.FindStringSubmatch(line)[1] ]
        // fmt.Println(first)
        // fmt.Println(last)
        total += first * 10 + last
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(total)
}
