package main

import (
	"fmt"
	"time"
)

const RANDOM_NUMBER_COUNT = 1_000_000

func main() {
	fmt.Printf("Cyrptographic random number generator")

	fileOperator := NewFileOperator()

	mersenneTwister := NewMersenneTwister(fileOperator, uint32(time.Now().UnixNano()), RANDOM_NUMBER_COUNT)
	mersenneTwister.Cyrpth(
		"MersenneTwister-1",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
	mersenneTwister.Cyrpth(
		"MersenneTwister-2",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)

	xorShift := NewXorShift(fileOperator, RANDOM_NUMBER_COUNT)
	xorShift.Cyrpth(
		"xorShift-1",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
	xorShift.Cyrpth(
		"xorShift-2",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)

	pcg := NewPCG(fileOperator, RANDOM_NUMBER_COUNT, 32)
	pcg.Cyrpth(
		"pcg-1",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
	pcg.Cyrpth(
		"pcg-2",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)

	blumBlumShub := NewBlumBlumShub(fileOperator, RANDOM_NUMBER_COUNT)
	blumBlumShub.Cyrpth(
		"blumBlumShub-1",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)
	blumBlumShub.Cyrpth(
		"blumBlumShub-2",
		"/Users/furkanmelih/personal_projects/cyrptographic-random-number-generator",
	)

}
