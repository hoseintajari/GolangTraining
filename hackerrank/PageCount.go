package hackerrank

func PageCount(n int32, p int32) int32 {
	front := p / 2

	back := (n - p) / 2

	if n%2 == 0 {
		back = (n + 1 - p) / 2
	}

	if front < back {
		return front
	} else {
		return back
	}
}
