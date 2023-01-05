package main

import (
    "bufio"
    "io"
    "os"
    "fmt"
)

func count(r io.Reader) int {
    scanner := bufio.NewScanner(r)
    scanner.Split(bufio.ScanWords)
    wc := 0

    for scanner.Scan() {
        wc++
    }
    return wc
}
func main() {

    fmt.Println(count(os.Stdin))
}
