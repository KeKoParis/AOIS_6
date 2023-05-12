package hash_table

import "math"

func Hash(key string) int32 {
	var hashNum int32

	k, p := 300, 1021 // polynomial constants
	for i, value := range key {
		hashNum += value * int32(math.Pow(float64(k), float64(i))) % int32(p)
	}

	return hashNum
}
