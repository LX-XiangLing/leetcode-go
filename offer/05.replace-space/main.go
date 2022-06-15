package _5_replace_space

func replaceSpace(s string) string {
	var res string
	for _, v := range s {
		if v == ' ' {
			res += "%20"
		} else {
			res += string(v)
		}
	}
	return res
}
