package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a.b/ac.go"
	fmt.Printf("%s\n", basename(str))
	fmt.Printf("%s\n", basename1(str))
}

func basename(str string) string {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' {
			str = str[i+1:]
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '.' {
			str = str[:i]
			break
		}
	}
	return str
}

func basename1(str string) string {
	i := strings.LastIndex(str, "/")
	str = str[i+1:]
	i = strings.LastIndex(str, ".")
	str = str[:i]
	return str
}
