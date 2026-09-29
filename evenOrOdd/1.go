package evenOrOdd

func IsEven(n int) bool {
	var isEven bool = false

	if n%2 == 0 {
		isEven = true
	}

	return isEven
}
