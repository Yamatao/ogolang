package main

import (
	"fmt"
)

const (
	barLength = 40
)

func PrintProgress(current, max float32) {
	fmt.Printf("\r")

	fmt.Printf("|")
	for i := 0; i < barLength; i++ {
		v := float32(i+1) / float32(barLength)
		if v <= current/max {
			fmt.Printf("=")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("| ")
}
