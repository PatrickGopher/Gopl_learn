package ChapterThree

func IsReverse(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	//记录每个字符出现的次数
	m := make(map[rune]int)
	n := make(map[rune]int)
	//以Unicode码作为map的key
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		n[v]++
	}
	for i, v := range m {
		if n[i] != v {
			return false
		}
	}
	return true
}
