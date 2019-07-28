package ChapterFour

import "sort"

// Prereqs记录了每个课程的前置课程
var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//将课程
func TopoSort(m map[string][]string) []string {
	var order []string
	//存储已经访问过的节点
	seen := make(map[string]bool)
	/**
	 *当匿名函数需要被递归调用时，我们必须首先声明一个变量（在上面的例子中，我们首先声明了 visitAll），
	 *再将匿名函数赋值给这个变量。
	 *如果不分成两部，函数字面量无法与visitAll绑定，我们也无法递归调用该匿名函数。
	**/
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
