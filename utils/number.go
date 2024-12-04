package utils

import "strconv"

func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}
