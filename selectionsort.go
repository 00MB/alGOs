package main

import "fmt"


func selection(array [7]int) [7]int {
	for i := 0; i < len(array); i++ {
		min_idx := i
		for j := i+1; j < len(array); j++ {
			if array[min_idx] > array[j] {
				min_idx = j
			}
		}
		array[i], array[min_idx] = array[min_idx], array[i]
	}
	return array
}

func main() {
	fmt.Println(selection([7]int{3, 8, 5, 4, 1, 9, -2}))
}