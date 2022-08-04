package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 50))
	// fmt.Println(SelectionSort([]int{5, 8, 1, 9, 2, 3, 7}))
	// fmt.Println(RecursiveArraySum([]int{2, 5, 6, 7, 10, 4}))
	// fmt.Println(CountNumbersInList([]int{2, 5, 6, 7, 10}))
	// fmt.Println(MaxNumberInList([]int{2, 5, 26, 7, 10}))
	slice := generateSlice(50)
	fmt.Println(QuickSort(slice))

}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
