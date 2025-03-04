package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Cyrptographic random number generator")
	mersenneTwister := NewMersenneTwister(uint32(time.Now().UnixNano()), 10)
	mersenneTwister.Cyrpth()
}
