package chapter4

import (
	"encoding/json"
	"fmt"
)

func UseComplexStruct() {
	//数组初始化
	var q [3]int = [3]int{1, 2, 3}
	ints := [...]int{1, 2}
	fmt.Println(q, ints)

	//长度也是数组的一种属性，所以不可以将[3]int{}与[4]int{}互相赋值
	// q:=[4]int{1,2}

	//初始化一个len为100的0数组，最后一位是-1，其余全是0
	i := [...]int{99: -1}
	fmt.Println(i)
}

func GetJson() {
	//JSON转换，字段需要大写,反射
	type Movie struct {
		Name  string
		Actor []string
		Year  int `json:"released,omitempty"` //released：json的key别名，omitempty：为0的时候不显示
	}
	var movies = []Movie{
		{Name: "green book", Actor: []string{"a", "b", "c"}, Year: 1998},
		{Name: "black book", Actor: []string{"c", "d", "e"}, Year: 7997},
	}
	marshal, _ := json.Marshal(movies)

	movie := Movie{Name: "green book", Actor: []string{"a", "b", "c"}, Year: 1998}
	bytes, _ := json.Marshal(movie)
	//格式化输出: jsonstring ,前缀，缩进
	indent, _ := json.MarshalIndent(movies, "", "  ")

	fmt.Printf("%s\n", marshal)
	fmt.Printf("%s\n", bytes)
	fmt.Printf("%s\n", indent)

	//反序列化
	//构造解析体
	type LittleMovie struct {
		Name string
	}
	//构建object
	var littleMovies []LittleMovie
	var littleMovie LittleMovie

	//传入jsonstring，和littleMovie的指针
	err2 := json.Unmarshal(bytes, &littleMovie)
	if err2 != nil {
		fmt.Println("parse json to object error")
		fmt.Println(err2.Error())
	}
	fmt.Println(littleMovie)
	//传入jsonstrings，和littleMovies的指针
	err := json.Unmarshal(marshal, &littleMovies)
	if err != nil {
		fmt.Println("parse jsons to objects error")
	}
	fmt.Println(littleMovies)

}
func GetGithubIssue() {
	const IssuesUrl = "https://api.github.com/search/issues"
	type IssusesSerahcResult struct {
		TotalCount int `json:"total_count"`
		Items      []*Issue
	}
	type Issue struct {
		Number int
	}
}
