package challenges

func CrossPattern(line int) string {
	retString := make([]byte, 0, line*(line+1))
	for i := 1; i <= line; i++ {
		for j := 1; j <= line; j++ {
			if j == i || j == line-(i-1) {
				retString = append(retString, ('*'))
			} else {
				retString = append(retString, (' '))
			}
		}
		retString = append(retString, ('\n'))

	}
	return string(retString)
}
