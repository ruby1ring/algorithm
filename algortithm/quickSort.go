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

func Heap(nums []int) {
	length := len(nums)

}
