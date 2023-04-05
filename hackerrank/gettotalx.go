package hackerrank

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var i, j int32
	var c int32
	for i = 1; i <= 100; i++ {
		factor := true

		//check first
		for j = 0; int(j) < len(a); j++ {
			if int32(i)%a[j] != 0 {
				factor = false
				break
			}
		}
		// check second
		if factor {
			for j = 0; int(j) < len(b); j++ {
				if b[j]%i != 0 {
					factor = false
					break
				}
			}
		}
		//update counter
		if factor {
			c++
		}

	}
	return c
}
