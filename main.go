package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Duration(1) * time.Second)

	for {
		select {
		case <-t.C:
			fmt.Println("hello world")
		}
	}
}
