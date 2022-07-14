package main

import (
	"math/rand"
	"time"
)

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers       = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // Number of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func valueGeneration(valType string, valLength int) string {
	var symbols string
	if valLength == 0 {
		valLength = int(src.Int63())
		for valLength > 66 {
			valLength /= 3
		}
		for valLength < 4 {
			valLength *= 3
		}
	}
	switch valType {
	case "num":
		symbols = numbers
	case "str":
		symbols = letters
	case "alp":
		symbols = letters + numbers
	}
	value := generator(valLength, symbols)
	return value
}

func generator(n int, symbols string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	server()
}
