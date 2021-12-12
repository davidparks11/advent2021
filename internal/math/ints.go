package math

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Sign(n int) int {
	if n == 0 {
		return 0
	} else if n < 0 {
		return -1
	}
	return 1
}