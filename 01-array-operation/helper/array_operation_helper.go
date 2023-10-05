package helper

func DuplicateAndShift(arr []int) []int {
	zeroCount := 0
	for _, num := range arr {
		if num == 0 {
			zeroCount++
		}
	}

	arrLength := len(arr)
	newIndex := arrLength + zeroCount - 1
	result := make([]int, arrLength)

	for i := arrLength - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if newIndex < arrLength {
				result[newIndex] = 0
			}
			newIndex--
		}

		if newIndex < arrLength {
			result[newIndex] = arr[i]
		}

		newIndex--
	}

	return result
}
