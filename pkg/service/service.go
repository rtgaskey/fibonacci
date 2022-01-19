package service

import (
	"strconv"
)

var currentIdx int = 0
var previous int = 0

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func GetCurrent() string {
	return strconv.Itoa(fibonacci(currentIdx))
}

func GetNext() string {
	next := strconv.Itoa(fibonacci(currentIdx + 1))
	currentIdx++

	return next
}

func GetPrevious() string {
	previous := "0"

	if currentIdx > 0 {
		previous = strconv.Itoa(fibonacci(currentIdx - 1))

		currentIdx--
	}

	return previous
}
