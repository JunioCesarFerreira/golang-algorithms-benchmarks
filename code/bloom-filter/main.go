package main

import (
	"encoding/binary"
)

// BloomFilter representa a estrutura do Bloom Filter
type BloomFilter struct {
	bitset    []uint64
	size      uint64
	hashCount uint64
}

// NewBloomFilter cria um novo Bloom Filter
func NewBloomFilter(size uint64, hashCount uint64) *BloomFilter {
	bitsetSize := (size + 63) / 64 // Tamanho em uint64
	return &BloomFilter{
		bitset:    make([]uint64, bitsetSize),
		size:      size,
		hashCount: hashCount,
	}
}

func (bf *BloomFilter) setBit(pos uint64) {
	bf.bitset[pos/64] |= (1 << (pos % 64))
}

func (bf *BloomFilter) getBit(pos uint64) bool {
	return (bf.bitset[pos/64] & (1 << (pos % 64))) != 0
}

// murmurHash3 implementa MurmurHash3 de 64 bits com semente
func murmurHash3(data []byte, seed uint64) uint64 {
	const m uint64 = 0xc6a4a7935bd1e995
	const r uint64 = 47
	length := uint64(len(data))
	hash := seed ^ (length * m)

	for len(data) >= 8 {
		k := binary.LittleEndian.Uint64(data[:8])
		k *= m
		k ^= k >> r
		k *= m

		hash ^= k
		hash *= m
		data = data[8:]
	}

	if len(data) > 0 {
		var tail uint64
		for i := len(data) - 1; i >= 0; i-- {
			tail <<= 8
			tail |= uint64(data[i])
		}
		hash ^= tail
		hash *= m
	}

	hash ^= hash >> r
	hash *= m
	hash ^= hash >> r

	return hash
}

// Add adiciona um item ao Bloom Filter
func (bf *BloomFilter) Add(item string) {
	for i := uint64(0); i < bf.hashCount; i++ {
		hashValue := murmurHash3([]byte(item), i) % bf.size
		bf.setBit(hashValue)
	}
}

// Contains verifica se um item pode estar no Bloom Filter
func (bf *BloomFilter) Contains(item string) bool {
	for i := uint64(0); i < bf.hashCount; i++ {
		hashValue := murmurHash3([]byte(item), i) % bf.size
		if !bf.getBit(hashValue) {
			return false
		}
	}
	return true
}

func main() {
	bf := NewBloomFilter(1000, 3)

	bf.Add("exemplo1")
	bf.Add("exemplo2")

	println(bf.Contains("exemplo1")) // true
	println(bf.Contains("exemplo2")) // true
	println(bf.Contains("exemplo3")) // false (pode ser true devido a falsos positivos)
}
