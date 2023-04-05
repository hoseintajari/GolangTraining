package hackerrank

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	yes := "YES"
	no := "NO"
	if x1 == x2 && v1 == v2 {
		return yes
	} else if x1 == x2 && v1 != v2 {
		return no
	} else {
		if v1 > v2 && ((x2-x1)%(v1-v2)) == 0 {
			return yes
		} else {
			return no
		}
	}
}
