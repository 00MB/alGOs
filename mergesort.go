package main

import "fmt"

func merge(array [7]int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = array[l + 1]
	}
	for j := 0; j < n2; j++ {
		R[j] = array[m + 1 + j]
	}


	fmt.Println(R, L)
	var i, j int
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			array[k] = L[i]
			i++
		} else {
			array[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		array[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		array[k] = R[j]
		j++
		k++
	}
	fmt.Println(array)
}

func mergeSort(array [7]int, l int, r int) {
	if l < r {
		m := l + (r - l)/2
		mergeSort(array, l, m)
		mergeSort(array, m + 1, r)
		merge(array, l, m, r)
	}

}

func main() {
	mergeSort([7]int{3, 8, 5, 4, 1, 9, -2}, 0, 6)
}