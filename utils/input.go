package utils

import (
	"bufio"
	"iter"
	"strings"
)

func Scan(input string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		var i int

		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			if !yield(i, scanner.Text()) {
				break
			}
			i++
		}
	}
}
