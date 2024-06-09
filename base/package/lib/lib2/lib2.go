package lib2

import "fmt"

// 当前lib2包 提供的 API
func Lib2Test() {
	fmt.Println("lib2Test() ...")
}

func init() {
	fmt.Println("lib2. init() ...")
}
