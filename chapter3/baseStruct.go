package chapter3

import (
	"fmt"
	"strconv"
)

//自增型常量
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Teusday
	Wednesday
	Thursday
	Friday
	Saturday
)

func UseBaseStruct() {
	//整数转换为int
	i, err := strconv.Atoi("123")
	//10进制，最长64位
	strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(i)
	}

	fmt.Println(Monday, Teusday, Wednesday, Thursday, Friday, Saturday, Sunday)

}
