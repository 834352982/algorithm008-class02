## 学习笔记

### 比较类排序

通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn)，因此也成为非线性时间比较类排序。

#### 初级排序 - O(n^2)

##### 选择排序（Selection_Sort）

每次找最小值，然后放到待排序数组的起始位置。

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015224719590-1433219824.gif)

```go
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
```

##### 插入排序（Insertion_Sort）

从前往后逐步构建有序序列；**对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入**。

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015225645277-1151100000.gif)

```go
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		needSort := arr[i]

		for preIndex >= 0 && arr[preIndex] > needSort {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}

		arr[preIndex+1] = needSort
	}
}
```

##### 冒泡排序（Bubble_Sort）

嵌套循环，每次查看相邻的元素，如果逆序，则交换。

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015223238449-2146169197.gif)

```go
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
```

#### 高级排序 - O(nlogn)

##### 快速排序（Quick_Sort）

数组取标杆pivot，将小元素放pivot左边，大元素放右边，然后依次对右边和右边的子数组继续快排；以达到整个序列有序。

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015230936371-1413523412.gif)

```go
func QuickSort(arr []int, begin int, end int) {
	if begin >= end {
		return
	}

	pivot := partition(arr, begin, end)
	QuickSort(arr, begin, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func partition(arr []int, begin int, end int) int {
	pivot := end
	// 小于pivot的元素的个数
	// 最后counter所在的位置就是pivot所在的位置
	counter := begin

	// 把所有小于pivot的元素都移到counter左边
	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}

	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}
```

##### 归并排序（Merge_Sort） - 分治

- 把长度为n的输入序列分为两个长度为n/2的子序列
- 对这两个子序列分别采用归并排序
- 将两个排序好的子序列合并成一个最终的排序序列

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015230557043-37375010.gif)

```go
func MergeSort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	mid := (right-left)>>1 + left

	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
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
```

 归并和快排具有相似性，但步骤顺序相反

- 归并：先排序左右子数组，然后合并两个有序子数组
- 快排：先分出左右子数组，然后对于左右子数组进行排序

##### 堆排序（Heap_Sort） - 堆插入O(logn)，取最大/小值O(1)

- 数组元素依次建立小顶堆
- 依次取堆顶元素，并删除

![](https://images2017.cnblogs.com/blog/849589/201710/849589-20171015231308699-356134237.gif)

```c++
void heap_sort(int a[], int len) {
    priority_queue<int,vector<int>,greater<int>> q;

    for(int i = 0; i < len; i ++) {
        q.push(a[i]);
	}

    for(int i = 0; i < len; i ++) {
        a[i] = q.pop();
    }
}
```

```go
func HeapSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, len int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < len && arr[left] > arr[largest] {
		largest = left
	}

	if right < len && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, len, largest)
	}
}
```
