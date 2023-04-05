package hackerrank

func BirthdayCakeCandles(candles []int32) int32 {
	max := candles[0]
	count := 0
	for _, v := range candles {
		if v > max {
			max = v
		}
	}
	for _, v := range candles {
		if max == v {
			count++
		}
	}
	return int32(count)
}
