package util

import (
	"fmt"
	"time"
)

func CreateCronJob() {
	fmt.Println(time.Now())
	t := time.NewTicker(time.Second * 30)
	for v := range t.C {
		fmt.Println(v, "dwqeqw")
	}
}
