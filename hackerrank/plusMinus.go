package hackerrank

import "fmt"

func PlusMinus(arr []int32) {
	// Write your code here
	var positive float32
	var negative float32
	var zeros float32
	for _, v := range arr {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else if v == 0 {
			zeros++
		}
	}
	retp := positive / float32(len(arr))
	retn := negative / float32(len(arr))
	retz := zeros / float32(len(arr))
	fmt.Println(retp)
	fmt.Println(retn)
	fmt.Println(retz)
}
