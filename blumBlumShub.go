package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

type BlumBlumShub struct {
	fileOperator      *FileOperator
	seedNumber        *big.Int
	randomNumberCount int
	p                 *big.Int
	q                 *big.Int
	M                 *big.Int
	X                 *big.Int
}

func NewBlumBlumShub(fileOperator *FileOperator, randomNumberCount int) *PCG {
	return &PCG{
		fileOperator:      fileOperator,
		randomNumberCount: randomNumberCount,
	}
}

func (b *BlumBlumShub) generateLargePrime(bits int) (*big.Int, error) {
	for {
		prime, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, err
		}

		if new(big.Int).Mod(prime, big.NewInt(4)).Cmp(big.NewInt(3)) == 0 {
			return prime, nil
		}
	}
}

func (b *BlumBlumShub) fillParameters() {
	p, err := b.generateLargePrime(64)
	if err != nil {
		panic(err)
	}

	q, err := b.generateLargePrime(64)
	if err != nil {
		panic(err)
	}

	for p.Cmp(q) == 0 {
		q, err = b.generateLargePrime(64)
		if err != nil {
			panic(err)
		}
	}

	b.q = q
	b.p = p

	b.M = new(big.Int).Mul(b.q, b.p)

	seed, err := rand.Int(rand.Reader, new(big.Int).Mul(b.q, b.p))
	if err != nil {
		panic(err)
	}
	b.seedNumber = seed

	seedModM := new(big.Int).Mod(seed, b.M)
	if new(big.Int).GCD(nil, nil, seedModM, b.M).Cmp(big.NewInt(1)) != 0 {
		panic(err)
	}

	b.X = new(big.Int).Exp(seed, big.NewInt(2), b.M)
}

func (p *BlumBlumShub) Cyrpth(fileName string, filePath string) {
	p.fillParameters()

	var strVals []string = make([]string, 0)

	for i := 0; i < p.randomNumberCount; i++ {
		number := p.extractNumber()
		fmt.Println(number)

		strVals = append(strVals, strconv.FormatUint(number, 10))
	}

	fileName = fileName + ".txt"
	p.fileOperator.SaveArrayToTxtFile(strVals, fileName, filePath)
}

func (b *BlumBlumShub) extractNumber() uint64 {
	b.X.Exp(b.X, big.NewInt(2), b.M)
	number := uint64(new(big.Int).Mod(b.X, big.NewInt(2)).Int64())

	var newNumber uint64 = 0
	for {
		if number == 0 {
			break
		}
		digit := newNumber % 10
		digit = digit % 2
		newNumber = newNumber*10 + digit

		number = number / 10
	}

	return newNumber
}
