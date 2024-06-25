package leetcode

func HeapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr)-1)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr, 0, i-1)
	}
}

func maxHeapify(arr []int, begin int, end int) {
	parent := begin
	child := 2*parent + 1
	for child <= end {
		if child+1 <= end && arr[child+1] > arr[child] {
			child++
		}
		if arr[child] < arr[parent] {
			return
		} else {
			arr[parent], arr[child] = arr[child], arr[parent]
			parent = child
			child = 2*parent + 1
		}
	}
}
