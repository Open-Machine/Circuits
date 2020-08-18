package utils

import (
	"math"
)

func IsOverflow(num uint, availableBits int) bool {
	largestNumber := math.Pow(2, float64(availableBits))
	return num >= uint(math.Floor(largestNumber))
}
