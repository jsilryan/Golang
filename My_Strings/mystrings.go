package mystrings

// It is to be used in the Hellogo package
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}

	return result
}
