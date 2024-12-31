package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5, 7, 7, 1, 2, 2, 2}
	maximum := array[0]

	for i := 0; i < len(array); i++ {
		if array[i] > maximum {
			maximum = array[i]
		}
	}

	fmt.Println(maximum)

	hash := make([]int, maximum+1, maximum+1)
	for i := 0; i < len(array); i++ {
		hash[array[i]] += 1
	}
	fmt.Print(hash)
}
