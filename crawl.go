// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
    "strings"
    "gopl.io/ch5/links"
)
var tokens = make(chan struct{}, 20)

var seen = make(map[string]bool)


func crawl(url string, depth int) {
	fmt.Println(url)
    tokens <- struct{}{} // acquire a token
	list:=links.Extract(url)
     <-tokens
    for i:=0; i<len(list)&&depth>0; i++{
        if !seen[list[i]]{
            crawl(list[i], depth-1)
        }
    }
    return
}



func main() {
    s:=""
    var n int
    fmt.Scanln(&s)
    fmt.Scanln(&n)
    crawl(s, n)
}
