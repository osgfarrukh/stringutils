package stringutils

func Reverse(s string) string {
	str := ""
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}
