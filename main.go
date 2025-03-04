package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Cyrptographic random number generator")

	fileOperator := NewFileOperator()
	mersenneTwister := NewMersenneTwister(fileOperator, uint32(time.Now().UnixNano()), 1_000_000)
	mersenneTwister.Cyrpth(
		"MersenneTwister-1",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
	mersenneTwister.Cyrpth(
		"MersenneTwister-2",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
}
