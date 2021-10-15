package algortithm

func Quicksort(nums []int, left, right int) {
	if left < right {
		l := left
		r := right
		key := nums[l]
		for l < r {
			for nums[r] > key {
				r--
			}
			if l < r {
				nums[l] = nums[r]
				l++
			}
			for nums[l] < key {
				l++
			}
			if l < r {
				nums[r] = nums[l]
				r--
			}
		}
		nums[l] = key
		Quicksort(nums, left, l-1)
		Quicksort(nums, l+1, right)
	}
}

//构建最小堆
func MinHeap(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		MinHeapSort(nums, i, length-1)
	}
	for i := length - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		MinHeapSort(nums, 0, i-1)
	}
}

func MinHeapSort(nums []int, index, bound int) {
	i := 2*index + 1
	if i > bound {
		return
	}
	n := 2*index + 2
	max := i
	if n <= bound && nums[n] > nums[i] {
		max = n
	}
	if nums[max] < nums[index] {
		return
	}
	nums[max], nums[index] = nums[index], nums[max]
	MinHeapSort(nums, max, bound)
}

//构建最大堆
func MaxHeap(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		MaxHeapSort(nums, i, length-1)
	}
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		MaxHeapSort(nums, 0, i-1)
	}
}

func MaxHeapSort(nums []int, index, bound int) {
	l := 2*index + 1
	if l > bound {
		return
	}
	r := 2*index + 2
	min := l
	if r <= bound && nums[l] > nums[r] {
		min = r
	}
	if nums[min] > nums[index] {
		return
	}
	nums[min], nums[index] = nums[index], nums[min]
	MaxHeapSort(nums, min, bound)
}
