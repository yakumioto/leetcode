package problem0001

func TwoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	hashNum := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if v, ok := hashNum[target-nums[i]]; ok {
			result = append(result, i, v)

			return result
		}

		hashNum[nums[i]] = i
	}

	return nil
}
