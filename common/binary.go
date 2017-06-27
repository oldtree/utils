package common

func CountBitNumber(value uint64) int {
	count := 0
	for ; value > 0; value = value >> 1 {
		count = count + int(value&1)
	}
	return count
}

func IsPowerOf2(value uint64) bool {
	return value&(value-1) == 0
}
