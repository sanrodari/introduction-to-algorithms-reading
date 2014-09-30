package main

import "fmt"

func main() {
	array := []int{5, 2, 4, 6, 1, 3}

	for j := 1; j < len(array); j++ {
		key := array[j]
		i := j - 1

		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i = i - 1
		}

		array[i+1] = key
	}

	fmt.Println("array == ", array)
}
