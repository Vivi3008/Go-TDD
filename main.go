package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())
	fmt.Printf("Hello World %v", isTimeToDebit())
}

func isTimeToDebit() bool {
	now := time.Now().UTC()
	return now.Hour() >= 13 && now.Hour() < 23
}
