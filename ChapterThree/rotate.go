package ChapterThree

func Rotate(s []int, i int) {
	ReverseSlice(s[:i])
	ReverseSlice(s[i:])
	ReverseSlice(s)
}

func Rotate2(s []int, i int) []int {
	slen := len(s)
	for _, v := range s {
		s = append(s, v)
	}
	i = i % slen
	news := s[i : slen+i]
	return news
}
