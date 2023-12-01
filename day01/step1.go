package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    regexfirst, _ := regexp.Compile("^.*?(\\d)")
    regexlast, _ := regexp.Compile("^.*(\\d)")
    total := 0

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        line := scanner.Text()
        // fmt.Println(scanner.Text())
        first, _ := strconv.Atoi( regexfirst.FindStringSubmatch(line)[1] )
        last, _  := strconv.Atoi( regexlast.FindStringSubmatch(line)[1] )
        total += first * 10 + last
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(total)
}
