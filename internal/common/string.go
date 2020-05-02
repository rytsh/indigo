package common

// TrimSuffixAll remove all end of character of string
func TrimSuffixAll(val string, ch byte) string {
	for val != "" && val[len(val)-1] == ch {
		val = val[:len(val)-1]
	}
	return val
}
