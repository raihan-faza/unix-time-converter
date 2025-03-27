package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	userInput, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	humanReadableTime := time.Unix(int64(userInput), 0)
	fmt.Printf("Human readable time: %s \n", humanReadableTime)
}
