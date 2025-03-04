package main

import "fmt"

// FORMULA
// STEP 1: INITIALIZATION
// s0 = seed
// si = (F . (si-1 XOR (si-1 >> (w - 2)))) + i mod 2^w

// w=32 (bit width of each number)
// F=1812433253 (a multiplier constant)
// ≫ represents right bit-shift
// i is the index in the state array

// STEP 2: TWISTING TRANSFORMATION
// y = (si&u) + (si+1 modn&1)
// si = s(i+m) mod n XOR (y >> 1)
// if y is a odd an additional transformation is applied
// si = si XOR a

// & = bitwise and operator
// u=2^31 (upper mask, 31 bits)
// l=2^31 - 1 (lower mask, 31 bits)
// m = 397 (offset index update)
// a=0x9908B0DF (bitwise constant for transformation)
// This transformation ensures that the state array is shuffled and prevents predictable sequences.

// STEP 3: Extracting Numbers
// y = s
// y = y XOR (y >> u)
// y = y XOR ((y << s) & b)
// y = y XOR ((y << t)& c)
// y = y XOR (y >> l)

const (
	n         = 624
	mC        = 397
	matrixA   = 0x9908B0DF // Constant for the twist transformation
	upperMask = 0x80000000 // Most significant bit (32-bit)
	lowerMask = 0x7FFFFFFF // Least significant bits
)

type MersenneTwister struct {
	mt                [n]uint32
	index             uint32
	seedNumber        uint32
	randomNumberCount int
}

func NewMersenneTwister(seedNumber uint32, randomNumberCount int) *MersenneTwister {
	return &MersenneTwister{seedNumber: seedNumber, randomNumberCount: randomNumberCount}
}

func (m *MersenneTwister) Cyrpth() {
	m.Seed()

	for i := 0; i < m.randomNumberCount; i++ {
		fmt.Println(m.ExtractNumber())
	}

}

func (m *MersenneTwister) Seed() {
	m.mt[0] = m.seedNumber
	for i := 1; i < n; i++ {
		m.mt[i] = 1812433253*(m.mt[i-1]^(m.mt[i-1]>>30)) + uint32(i)

	}
	m.index = n
}

func (m *MersenneTwister) Twist() {
	for i := 0; i < n; i++ {
		y := (m.mt[i] & upperMask) | (m.mt[(i+1)%n] & lowerMask)
		m.mt[i] = m.mt[(i+mC)%n] ^ (y >> 1)
		if y%2 != 0 {
			m.mt[i] ^= matrixA
		}
	}
	m.index = 0
}

func (m *MersenneTwister) ExtractNumber() uint32 {
	if m.index >= n {
		m.Twist()
	}

	y := m.mt[m.index]
	m.index++

	y ^= (y >> 11)
	y ^= (y << 7) & 0x9D2C5680
	y ^= (y << 15) & 0xEFC60000
	y ^= (y >> 18)

	var newY uint32 = 0
	for {
		if y == 0 {
			break
		}
		digit := y % 10
		digit = digit % 2
		newY = newY*10 + digit

		y = y / 10
	}

	return newY
}
