package hackerrank

import "fmt"

func Staircase(n int32) {
	var k int32
	for i := 1; int32(i) <= int32(n); i++ {
		for j := i; int32(j) < n; j++ {
			fmt.Print(" ")
		}
		for k = n - int32(i); int32(k) < n; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
