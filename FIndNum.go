package main

import "fmt"

func main() {
	var arr = [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	fmt.Println(FIndNum(arr, 20))
}

func FIndNum(arr [][]int, num int) bool {
	i := len(arr) - 1
	j := 0

	for i > 0 && j < len(arr[i]) {
		if arr[i][j] > num {
			i--
		} else if arr[i][j] < num {
			j++
		} else {
			return true
		}
	}

	return false
}
