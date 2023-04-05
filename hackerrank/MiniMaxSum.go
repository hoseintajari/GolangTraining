package hackerrank

import (
	"fmt"
)

func MiniMaxSum(arr []int32) {
	var max int32
	var min = arr[0]
	var summax int
	var summin int
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}

	}
	fmt.Println(max, min)
	for _, v := range arr {
		if v != min {
			summax += int(v)
		}
		if v != max {
			summin += int(v)
		}
	}
	fmt.Println(summin, summax)

}
