package main

import (
	"./ChapterThree"
	"fmt"
)

//基础类型：数字、字符、布尔
//复合类型：数组、结构体
//引用类型：指针、切片、字典、函数、通道
//接口类型

func main() {
	fmt.Println(ChapterThree.IsReverse("aavcx", "xcava"))
	//new返回指针
	p := new(int)
	fmt.Println(*p)
	*p = 5
	fmt.Println(*p)
	//最大公倍数
	fmt.Println(gcd(6, 8))
	//科学计数法
	a := 1e2
	fmt.Println(int(a))

	//六进制
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	//处理中文字符串rune
	str := "Hello,世界"
	fmt.Println(str)
	fmt.Printf("% x\n", str)
	x := []rune(str)
	fmt.Printf("%x\n", x)
	fmt.Println(string(x))
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\t\n", i, r, r)
	}
	fmt.Printf("%q", str)
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
