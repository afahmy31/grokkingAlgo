package main

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	less, greater = QuickSort(less), QuickSort(greater)

	less = append(less, pivot)
	less = append(less, greater...)
	return less
}
