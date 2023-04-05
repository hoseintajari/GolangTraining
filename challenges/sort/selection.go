package sort

func Selection(array []int) []int {
	var index int
	var temp int
	for i := 0; i < len(array)-1; i++ {
		index = i
		// Find index of minimum element
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
		}
		temp = array[i]
		array[i] = array[index]
		array[index] = temp
	}
	return array
}
