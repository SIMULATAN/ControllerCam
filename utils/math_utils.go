package utils

import (
	"math"
)

func CalculateExponentialValue(value float64, factor float64) float64 {
	return math.Pow(value, factor)
}
