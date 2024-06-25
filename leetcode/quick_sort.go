package leetcode

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, begin int, end int) {
	if end <= begin {
		return
	}
	pivotIndex := partition(arr, begin, end)
	quickSort(arr, begin, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end)
}

func partition(arr []int, begin int, end int) int {
	pivot := arr[begin]
	for begin < end {
		for begin < end && arr[end] > pivot {
			end--
		}
		arr[begin] = arr[end]
		for begin < end && arr[begin] <= pivot {
			begin++
		}
		arr[end] = arr[begin]
	}
	arr[begin] = pivot
	return begin
}
