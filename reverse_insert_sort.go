package main

import "fmt"

func main() {
  array := []int{5, 2, 4, 6, 1, 3}

  for j := len(array) - 2; j >= 0; j-- {
    key := array[j]
    i := j + 1

    for i < len(array) && array[i] > key {
      array[i-1] = array[i]
      i = i + 1
    }

    array[i-1] = key
  }

  fmt.Println("array == ", array)
}
