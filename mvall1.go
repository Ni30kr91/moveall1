package main

import "fmt"

func main() {
	var arr = []int{5, 1, 33, 23, 1, -5, 1}
	fmt.Println(arr)
	var result = nitish(arr)
	fmt.Println(result)

}

func nitish(array []int) []int {
	result := []int{}

	onearray := []int{}

	for i := 0; i < len(array); i++ {
		if array[i] != 1 {
			result = append(result, array[i])
		} else {
			onearray = append(onearray, 1)
		}
	}

	result = append(result, onearray...)
	return result
}
