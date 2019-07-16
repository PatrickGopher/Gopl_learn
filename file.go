package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		f.WriteString(buf)
	}
}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r:= bufio.NewReader(f)

	for{
		buf, err:= r.ReadBytes('\n')
		if err!=nil{
			if err == io.EOF{
				break
			}
			fmt.Println(err)
		}
		fmt.Print(string(buf))
	}
}

func main() {
	path := "./demo.txt"
	writeFile(path)
	readFile(path)
	readFileLine(path)
}
