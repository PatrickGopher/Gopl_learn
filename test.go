package main

//基础类型：数字、字符、布尔
//复合类型：数组、结构体
//引用类型：指针、切片、字典、函数、通道
//接口类型
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

func main() {
	//fmt.Println(ChapterThree.IsReverse("aavcx", "xcava"))
	////new返回指针
	//p := new(int)
	//fmt.Println(*p)
	//*p = 5
	//fmt.Println(*p)
	////最大公倍数
	//fmt.Println(gcd(6, 8))
	////科学计数法
	//a := 1e2
	//fmt.Println(int(a))
	//
	////六进制
	//o := 0666
	//fmt.Printf("%d %[1]o %#[1]o\n", o)
	//
	////处理中文字符串rune
	//str := "Hello,世界"
	//fmt.Println(str)
	//fmt.Printf("% x\n", str)
	//x := []rune(str)
	//fmt.Printf("%x\n", x)
	//fmt.Println(string(x))
	//for i, r := range str {
	//	fmt.Printf("%d\t%q\t%d\t\n", i, r, r)
	//}
	//fmt.Printf("%q", str)

	//const noDelay time.Duration = 0
	//const timeout = 5 * time.Minute
	//fmt.Printf("%T %[1]v\n", noDelay)
	//fmt.Printf("%T %[1]v\n", timeout)

	//fmt.Printf("%d", KB)
	//a := []string{"a", "ab"}
	//b := []string{"a", "ab"}
	//fmt.Println(equal(a, b))

	//var runes []rune
	//for _, v := range "Hello,世界" {
	//	runes = append(runes, v)
	//}
	//fmt.Printf("%q\n", runes)

	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d %d\t%v\n", i, cap(y), y)
	//	x = y
	//}
	//
	//arr := []int{1, 8, 5, 2, 3}
	//ChapterThree.Reverse(&arr)
	//fmt.Println(arr)
	//ChapterThree.Rotate(arr, 2)
	//fmt.Println(arr)
	//arr1 := ChapterThree.Rotate2(arr, 2)
	//fmt.Println(arr1)

	//ages := map[string]int{
	//	"bbbba": 2,
	//	"aaaab": 1,
	//	"aaaaa": 1,
	//}
	//ages["v"] = 1
	//ages["a"] = 2
	//delete(ages, "v")
	//ages["v"] = ages["v"] + 1
	//fmt.Println(ages)
	//ChapterThree.SortMap(ages)
}

func appendInt(x []int, i int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen < cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = i
	return z
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
