package main

import (
	"fmt"
)

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 5
	fmt.Println(*p)
	//fmt.Println(gcd(6, 8)
	a := 1e2
	fmt.Println(int(a))

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	fmt.Println(HasPrefix("aaab", "aaa"))
	fmt.Println(Contains("aaabsd", "abs"))

	str := "Hello,世界"
	fmt.Println(str)
	fmt.Printf("% x\n", str)
	x := []rune(str)
	fmt.Printf("%x\n", x)
	fmt.Println(string(x))
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\t\n", i, r, r)
	}
	fmt.Printf("%q",str)
}

//func gcd(x int, y int) int {
//	for y != 0 {
//		x, y = y, x%y
//	}
//	return x
//}
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
