package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
