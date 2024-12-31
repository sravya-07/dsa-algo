package main
import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 6, 7}
	maximum := array[0]
	length := len(array)

	for i := 0; i < length; i++ {
		if array[i] > maximum {
			maximum = array[i]
		}
	}

	hash := make(map[int]int)
	for i := 0; i < length; i++ {
		_, present := hash[array[i]]
		if present {
			hash[array[i]] += 1
		} else {
			hash[array[i]] += 1
		}
	}
	fmt.Println(hash)
}
