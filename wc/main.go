package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
    if(countLines){
        scanner.Split(bufio.ScanLines)
    } else if(countBytes) {
        scanner.Split(bufio.ScanBytes)
    } else {
        scanner.Split(bufio.ScanWords)
    }

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc

}

func main() {
    lines := flag.Bool("l", false, "Count Lines")
    nbBytes := flag.Bool("b", false, "Count Bytes")
    flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *nbBytes))
}
