package kata

func GetMiddle(s string) string {
	if len(s)%2 != 0 {
		return string(s[len(s)/2])
	}
	return s[len(s)/2-1 : len(s)/2+1]
}
