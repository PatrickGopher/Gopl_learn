package ChapterThree

func Reverse(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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
