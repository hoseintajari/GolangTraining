package hackerrank

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	houseDistanceStart := s
	houseDistanceEnd := t
	appleTreePosition := a
	orangeTreePosition := b

	var _apples []int32
	var _oranges []int32

	for i := 0; i < len(apples); i++ {
		sum := appleTreePosition + apples[i]

		if sum >= houseDistanceStart && sum <= houseDistanceEnd {
			_apples = append(_apples, sum)
		}
	}

	for i := 0; i < len(oranges); i++ {
		sum := orangeTreePosition + oranges[i]

		if sum >= houseDistanceStart && sum <= houseDistanceEnd {
			_oranges = append(_oranges, sum)
		}
	}

}
