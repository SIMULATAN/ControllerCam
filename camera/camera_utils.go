package camera

/**
 * Returns the `num`, unless it is less than `minDistance` from zero, in which case it returns `minDistance` with the sign of `num`.
 */
func signedMinDistanceFromZero(num, minDistance int16) int16 {
	if num > 0 && num < minDistance {
		return minDistance
	}
	if num < 0 && num > -minDistance {
		return -minDistance
	}
	return num
}
