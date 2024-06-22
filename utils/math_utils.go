package utils

import (
	"math"
)

func CalculateExponentialValue(value float64, factor float64) float64 {
	// sign is either 1 or -1
	// this allows us to calculate the exponential value of a negative number
	sign := value / math.Abs(value)
	return sign * math.Pow(math.Abs(value), factor)
}
