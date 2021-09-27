package main

import (
	"fmt"
	"time"
)

func Fn[T any](t T) {
	fmt.Println(t)
}


func main() {
	Fn(0)
	Fn(0.0)
	Fn(int64(0))
	Fn(uint64(0))
	Fn("0")
	Fn([]byte{'0'})
	Fn(struct{}{})
	Fn(time.Now())
}
