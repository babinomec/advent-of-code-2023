package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "regexp"
    "strconv"
)

func main() {


    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    rgame, _ := regexp.Compile("^Game (\\d+)")
    rdraw, _ := regexp.Compile("(\\d+) (red|blue|green)")
     total := 0

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        split := strings.Split( line, ":")
        game, _ := strconv.Atoi( rgame.FindStringSubmatch( split[0] )[1] )
        draws := strings.Split( split[1], ";" )

        min := map[string]int{ "red": 0, "green": 0, "blue": 0 }

        for _, draw := range draws {
            matches := rdraw.FindAllStringSubmatch(draw, -1)
            // fmt.Println(matches)

            for _, subdraw := range matches {
                number, _ := strconv.Atoi(subdraw[1])
                color := subdraw[2]
                if ( number > min[ color ] ) {
                    min[color] = number
                } 
            }
        }

        power := 1
        for _, v := range min {
            power *= v
        }
        fmt.Println( "Game", game, "power", power )
        total += power
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(total)
}
