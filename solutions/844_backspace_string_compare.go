package solutions

func backspaceString(s string) string {
	index := 0
	for index < len(s) {
		if s[index] == '#' && len(s) > 0 {
			if index == 0 {
				s = s[index+1:]
			} else {
				s = s[:index-1] + s[index+1:]
				index--
			}
		} else {
			index++
		}
	}
	return s
}

func BackspaceCompare(s string, t string) bool {
	return backspaceString(s) == backspaceString(t)
}
