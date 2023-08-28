// Author: Michael Bopp
// Filename: goroutine.go
// Description: Go program go routines & channels.
// Date Created: 8/28/23
// Last Edited: 8/28/23

package main

import (
	"fmt"
	"time"
)

func say(text string, ch chan string) {
	time.Sleep(time.Second)
	fmt.Println(text)
	ch <- "Done"
}

func main() {
	ch := make(chan string)
	go say("Hello from goroutine!", ch)
	fmt.Println("Hello from main!")
	<-ch
}
