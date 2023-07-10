package defusethebomb

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))

	if len(code) == 0 {
		return result
	}

	for i := 0; i < len(code); i++ {
		sum := 0
		if k > 0 {
			start := i + 1
			for j := 0; j < k; j++ {
				if start > len(code)-1 {
					start = 0
				}
				sum += code[start]
				start++
			}

			result[i] = sum
		} else {
			start := i - 1
			for j := k; j < 0; j++ {
				if start < 0 {
					start = len(code) - 1
				}
				sum += code[start]
				start--
			}

			result[i] = sum
		}
	}

	return result
}
