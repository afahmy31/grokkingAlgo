package main

func RecursiveArraySum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	last := len(arr) - 1
	return arr[last] + RecursiveArraySum(arr[:last])
}
