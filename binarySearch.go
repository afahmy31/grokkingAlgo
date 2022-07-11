package main

func BinarySearch(array []int, number int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + high/2
		guess := array[mid]
		if guess == number {
			return mid
		}
		if guess > number {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
