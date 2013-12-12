package main

import (
	"fmt"
	"math/rand"
	"time"
)



func flip() bool {
	rand.Seed(time.Now().Unix())
	if(rand.Intn(2) == 0) {
		return true
	} else {
		return false
	}
}

func main() {

	isHeads := flip()

	if(isHeads) {
		fmt.Printf("Heads.")
	} else {
		fmt.Printf("Tails.")
	}

    
}