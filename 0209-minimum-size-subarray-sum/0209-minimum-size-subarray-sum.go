func minSubArrayLen(tar int, nums []int) int {
	start := 0
	end := 0
	res := math.MaxInt
	sum := 0
	for end < len(nums) {
		sum += nums[end]

		for sum >= tar {
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}

		end++
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}