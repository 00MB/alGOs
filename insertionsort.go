package main

import (
	"fmt"
)

func insertion(array [7]int) ([7]int){
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && key < array[j] {
			array[j + 1] = array[j]
			j--
		}
		array[j + 1] = key
	}
	return array
}

func main() {
	fmt.Println(insertion([7]int{3, 8, 5, 4, 1, 9, -2}))
}