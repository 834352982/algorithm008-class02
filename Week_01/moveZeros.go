func moveZeroes(nums []int)  {
    snowBollSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowBollSize++
		} else if snowBollSize > 0 {
			nums[i-snowBollSize] = nums[i]
			nums[i] = 0
		}
	}
}