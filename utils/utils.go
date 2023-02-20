package utils

// Abs returns the absolute value for an int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
