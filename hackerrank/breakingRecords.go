package hackerrank

func BreakingRecords(scores []int32) []int32 {
	ret := make([]int32, 2, 2)

	highScore := scores[0]
	lowScore := scores[0]
	for _, v := range scores {
		if v > highScore {
			highScore = v
			ret[0]++
		} else if v < lowScore {
			lowScore = v
			ret[1]++
		}

	}
	return ret
}
