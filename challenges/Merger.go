package challenges

func Merger(input1, input2 []int) []int {
	ret := make([]int, 0, len(input1)+len(input2))
	var a int
	var b int

	for i := 0; i < cap(ret)-1; i++ {
		if input1[a] < input2[b] {
			ret = append(ret, input1[a])
			a++
		} else {
			ret = append(ret, input2[b])
			b++
		}

	}
	return ret
}
