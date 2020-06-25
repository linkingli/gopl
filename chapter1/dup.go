package chapter1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
//输出命令行参数
func PrintArgs() {
	//method1
	s, sep := "", ""
	for _, args := range os.Args[:] {
		s += sep + args
		sep = ""
	}
	fmt.Println(s)
	//method2
	fmt.Println(strings.Join(os.Args[:], ""))
}
//统计文本重复行，标准输入读取，go<test.txt
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println(input)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

//统计文本重复行，文件读取
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			open, err := os.Open(file)
			if err != nil {
				fmt.Println("open file error")
			}
			countLines(open, counts)
			open.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

func countLines(stdin *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
//统计文本重复行，文件一次读取，go test.txt
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//一次读取file中所有数据
		data, err := ioutil.ReadFile(filename)
		fmt.Println(string(data))
		if err != nil {
			fmt.Println("open file error")
			continue
		}
		//将读取出来的数据以分行符分割
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}