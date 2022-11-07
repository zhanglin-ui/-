package main

import "fmt"

func main() {
	var arr = []int{3, 2, 6, 4, 5, 1, 2}
	fmt.Println(RepeatNum(arr))
}

func RepeatNum(arr []int) int {
	if len(arr) <= 0 {
		return -1
	}

	start := 1
	end := len(arr) - 1
	for end >= start {
		mid := ((end - start) >> 1) + start
		count := CalCount(arr, start, mid)
		if start == end {
			if count > 1 {
				return start
			} else {
				return -1
			}
		}

		if count > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func CalCount(arr []int, start int, end int) (count int) {
	count = 0
	if len(arr) <= 0 || start <= 0 || end >= len(arr) {
		return -1
	}

	for _, num := range arr {
		if num >= start && num <= end {
			count++
		}
	}
	return count
}
