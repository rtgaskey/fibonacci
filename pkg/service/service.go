package service

var currentIdx int = 0

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func GetCurrent() int {
	// return strconv.Itoa(fibonacci(currentIdx))
	return fibonacci(currentIdx)
}

func GetNext() int {
	// next := strconv.Itoa(fibonacci(currentIdx + 1))
	next := fibonacci(currentIdx + 1)

	currentIdx++

	return next
}

func GetPrevious() int {
	previous := 0

	if currentIdx > 0 {
		// previous = strconv.Itoa(fibonacci(currentIdx - 1))
		previous = fibonacci(currentIdx - 1)

		currentIdx--
	}

	return previous
}
