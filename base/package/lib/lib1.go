package lib

import "fmt"

// 当前lib1包 提供的 API
func Lib1Test() {
	fmt.Println("lib1Test() ...")
}

func init() {
	fmt.Println("lib1. init() ...")
}
