package main

import (
	"fmt"
	"strconv"
)

/*
FORMULA

X(n + 1) = aXn + c mod m

Xn: the internal state of the generator at step
a = multiplier (typically a large odd number)
c = increment (odd number, ensures full period)
m = modulus (often 2^64 or 2^128)
X(n+1) = next internal state

Yn = bitwise_shuffle(Xn)
*/

const (
	pcgMultiplier = 6364136223846793005
	pcgIncrement  = 1442695040888963407
)

type PCG struct {
	fileOperator      *FileOperator
	seedNumber        uint64
	randomNumberCount int
}

func NewPCG(fileOperator *FileOperator, randomNumberCount int, seedNumber uint64) *PCG {
	return &PCG{
		fileOperator:      fileOperator,
		seedNumber:        seedNumber,
		randomNumberCount: randomNumberCount,
	}
}

func (p *PCG) Cyrpth(fileName string, filePath string) {
	var strVals []string = make([]string, 0)

	for i := 0; i < p.randomNumberCount; i++ {
		number := p.extractNumber()
		fmt.Println(number)

		strVals = append(strVals, strconv.FormatUint(uint64(number), 10))
	}

	fileName = fileName + ".txt"
	p.fileOperator.SaveArrayToTxtFile(strVals, fileName, filePath)
}

func (p *PCG) extractNumber() uint64 {
	p.seedNumber = (p.seedNumber*pcgMultiplier + pcgIncrement) & ((1 << 64) - 1)

	xorShifted := ((p.seedNumber >> 18) ^ p.seedNumber) >> 27
	rot := p.seedNumber >> 59

	result := (xorShifted >> rot) | (xorShifted<<((-rot)&31))&((1<<32)-1)
	var newResult uint64 = 0
	for {
		if result == 0 {
			break
		}
		digit := result % 10
		digit = digit % 2
		newResult = newResult*10 + digit

		result = result / 10
	}
	return newResult
}
