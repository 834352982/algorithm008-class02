func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(arr []int, left int, right int) int {
	if left >= right {
		return 0
	}

	mid := (right-left)>>1 + left

	count := mergeSort(arr, left, mid) + mergeSort(arr, mid+1, right)

    for i, j := left, mid + 1; i <= mid; i ++ {
        for j <= right && arr[i] > arr[j]*2 {
            j ++
        }
        count += j - (mid + 1)
    }

    merge(arr, left, mid, right)
    return count
}

func merge(arr []int, left int, mid int, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	for p := 0; p < len(temp); p++ {
		arr[left+p] = temp[p]
	}
}
