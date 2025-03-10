package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Xorshift generators use a series of XOR (exclusive OR) and bit shifts to update an internal state and generate random numbers. The key operations involve:

XOR-ing the current state with a bit-shifted version of itself.
Shifting bits left or right by a carefully chosen amount.
Repeating the process multiple times to ensure good randomness.

FORMULA
x = x XOR (x << a)
x = x XOR (x >> b)
x = x XOR (x << c)
*/

const (
	a = 13
	b = 17
	c = 5
)

type XorShift struct {
	xs                []uint32
	index             uint32
	randomNumberCount int
	fileOperator      *FileOperator
}

func NewXorShift(fileOperator *FileOperator, randomNumberCount int) *XorShift {
	return &XorShift{
		xs:                make([]uint32, randomNumberCount),
		fileOperator:      fileOperator,
		randomNumberCount: randomNumberCount,
	}
}

func (x *XorShift) Cyrpth(fileName string, filePath string) {
	x.seed()
	var strVals []string = make([]string, 0)
	for i := 0; i < x.randomNumberCount; i++ {
		number := x.extractNumber()
		fmt.Println(number)

		strVals = append(strVals, strconv.FormatUint(uint64(number), 10))
	}

	fileName = fileName + ".txt"
	x.fileOperator.SaveArrayToTxtFile(strVals, fileName, filePath)
}

func (x *XorShift) seed() {
	for i := 0; i < x.randomNumberCount; i++ {
		x.xs[i] = uint32(time.Now().UnixNano())
	}
	x.index = 0
}

func (x *XorShift) extractNumber() uint32 {
	num := x.xs[x.index]
	x.index++

	num ^= num << a
	num ^= num >> b
	num ^= num << c

	var newNum uint32 = 0
	for {
		if num == 0 {
			break
		}
		digit := num % 10
		digit = digit % 2
		newNum = newNum*10 + digit

		num = num / 10
	}

	return newNum
}
