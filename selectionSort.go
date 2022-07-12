package main

func smallestNumber(array []int) int {
	smallest := array[0]
	smallestIndex := 0
	for i, el := range array {
		if el < smallest {
			smallest = el
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SelectionSort(array []int) []int {
	var newArr []int
	for range array {
		smallestIndex := smallestNumber(array)
		newArr = append(newArr, array[smallestIndex])
		array = append(array[:smallestIndex], array[smallestIndex+1:]...)
	}
	return newArr
}
