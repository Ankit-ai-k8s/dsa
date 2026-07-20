func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	fmt.Println(nums)
	n := len(nums)
	res := [][]int{}
	for m := 0; m < n-2; m++ {
		if m > 0 && nums[m] == nums[m-1] {
			continue
		}
		start := m + 1
		end := n - 1
		for start < end {
			sum := nums[start] + nums[m] + nums[end]
			if sum == 0 {
				res = append(res, []int{nums[start], nums[m], nums[end]})
				fmt.Println(m, start, end)
				for start < n-1 && nums[start] == nums[start+1] {
					start++
				}
				for end > 0 && nums[end] == nums[end-1] {
					end--
				}
			}
			if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return res
}