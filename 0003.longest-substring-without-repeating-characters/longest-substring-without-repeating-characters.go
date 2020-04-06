package problem0003

func LengthOfLongestSubstringWithExhaustive(s string) int {
	charSet := map[byte]bool{}
	result := 0
	length := 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if ok := charSet[s[j]]; !ok {
				length++
				charSet[s[j]] = true
			} else {
				if length > result {
					result = length
				}

				charSet = map[byte]bool{}
				charSet[s[j]] = true
				length = 1
			}
		}

		if length > result {
			result = length
		}

		length = 0
		charSet = map[byte]bool{}
	}

	return result
}

func LengthOfLongestSubstringWithSlidingWindow(s string) int {
	charSet := map[byte]int{}
	left := 0
	result := 0

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < len(s); i++ {
		if pos, ok := charSet[s[i]]; ok {
			left = max(left, pos)
		}

		result = max(result, i-left+1)
		charSet[s[i]] = i + 1
	}

	return result
}
