package main

import "fmt"
import "time"

func main() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(200 * time.Millisecond)
		}
	}
}
