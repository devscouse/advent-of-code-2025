package core

import (
	"fmt"
	"log"
)

type BitArray struct {
	bits []uint64
}

func NewBitArray(size int) *BitArray {
	numWords := (size + 63) / 64
	log.Printf("Creating BitArray with %d words (%d bits)\n", numWords, numWords*64)
	return &BitArray{bits: make([]uint64, numWords)}
}

func (b *BitArray) Set(idx int) {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		newBits := make([]uint64, word+1)
		copy(newBits, b.bits)
		b.bits = newBits
		log.Printf("Expanding BitArray to %d words (%d bits)\n", len(b.bits), len(b.bits)*64)
	}
	b.bits[word] |= (1 << wordIdx)
}

func (b *BitArray) Unset(idx int) {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		return
	}
	b.bits[word] &^= (1 << wordIdx)
}

func (b *BitArray) IsSet(idx int) bool {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		panic(fmt.Errorf("bounds exceeded on BitArray: idx %d evaluates to word %d but BitArray only has %d words", idx, word, len(b.bits)))
	}
	return b.bits[word]&(1<<wordIdx) != 0
}
