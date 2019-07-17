//数字处理1000=》1,000
package ChapterThree

import (
	"bytes"
	"strings"
)

//func main() {
//	num := +12232131
//	fmt.Printf("%s\n", comma(strconv.Itoa(num)))
//	fmt.Printf("%s\n", comma1(strconv.Itoa(num)))
//	fmt.Printf("%s\n", comma2(strconv.Itoa(num)))
//}

func comma2(s string) string {
	var buffer bytes.Buffer
	if s[0] == '-' || s[0] == '+' {
		temp := s[0]
		s = s[1:]
		buffer.WriteByte(temp)
	}
	//分离整数与小数
	arr := strings.Split(s, ".")
	s = arr[0]
	l := len(s)

	for i := 0; i < l; i++ {
		buffer.WriteString(string(s[i]))
		if (i+1)%3 == l%3 && i != l-1 {
			buffer.WriteString(",")
		}
	}
	if len(arr) > 1 {
		buffer.WriteString(".")
		buffer.WriteString(arr[1])
	}
	return buffer.String()
}

func comma(s string) string {
	len := len(s)
	if len <= 3 {
		return s
	}
	return comma(s[:len-3]) + "," + s[len-3:]
}

func comma1(s string) string {
	var buffer bytes.Buffer
	l := len(s)
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))
		if (i+1)%3 == l%3 {
			buffer.WriteString(",")
		}
	}
	s = buffer.String()
	return s
}
