func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)

	for i := 0; i < len(nums); i ++{
		j := target - nums[i]
		if _, ok := mapNums[j]; ok {
            return []int{mapNums[j], i}
		}
		// 保证不会重复利用
		mapNums[nums[i]] = i
	}

	return []int{-1, -1}
}