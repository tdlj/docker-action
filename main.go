package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		t := time.NewTicker(time.Duration(1) * time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println("hello world")
			}
		}
	}()

	time.Sleep(5 * time.Second)

}
