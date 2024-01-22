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

    max := map[string]int{ "red": 12, "green": 13, "blue": 14 }

    file, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    rgame, _ := regexp.Compile("^Game (\\d+)")
    rdraw, _ := regexp.Compile("(\\d+) (red|blue|green)")
     total := 0

    scanner := bufio.NewScanner(file)

    game:
    for scanner.Scan() {
        line := scanner.Text()

        split := strings.Split( line, ":")
        game, _ := strconv.Atoi( rgame.FindStringSubmatch( split[0] )[1] )
        draws := strings.Split( split[1], ";" )

        for _, draw := range draws {
            matches := rdraw.FindAllStringSubmatch(draw, -1)
            // fmt.Println(matches)

            for _, subdraw := range matches {
                number, _ := strconv.Atoi(subdraw[1])
                color := subdraw[2]
                if ( number > max[ color ] ) {
                    fmt.Println( "Game", game, "not possible:", number, color )
                    continue game
                } 
            }
        }

        fmt.Println( "Game", game, "possible" )
        total += game
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(total)
}
