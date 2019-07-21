package ChapterThree

import (
	"fmt"
	"sort"
)

func SortMap(m map[string]int) {
	temp := make([]string, 0, len(m))
	for key := range m {
		temp = append(temp, key)
	}
	sort.Strings(temp)
	for _, name := range temp {
		fmt.Printf("%s\t%d\n", name, m[name])
	}
}
