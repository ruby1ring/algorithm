package main

import (
	quicksort "al/algortithm"
	"fmt"
)

func main() {
	arr := []int{3, 2, 1, 4, 2, 1, 4, 5}
	quicksort.MaxHeap(arr)
	fmt.Println(arr)
}
