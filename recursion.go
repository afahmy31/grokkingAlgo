package main

//ex4.1
func RecursiveArraySum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	last := len(arr) - 1
	return arr[last] + RecursiveArraySum(arr[:last])
}

//ex4.2
func CountNumbersInList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + CountNumbersInList(arr[1:])
}

//ex4.3
func MaxNumberInList(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}
	submMax := MaxNumberInList(arr[1:])
	if arr[0] > submMax {
		return arr[0]
	} else {
		return submMax
	}
}
