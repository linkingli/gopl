package chapter4

import "fmt"

func UseComplexStruct() {
	//数组初始化
	var q [3]int = [3]int{1, 2, 3}
	ints := [...]int{1, 2}
	fmt.Println(q, ints)

	//长度也是数组的一种属性，所以不可以将[3]int{}与[4]int{}互相赋值
	// q:=[4]int{1,2}
}
