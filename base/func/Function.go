package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return c
}

// 返回多个返回值，匿名的
func fool2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 666, 777
}

// 返回多个返回值，有形参名称的
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("------ fool3 -------")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	r1 = 10000
	r2 = 20000
	return
}

// func fool4(a string, b int) (r1 int, r2 int) {
// 	fmt.Println("------ fool4 -------")
// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)

// 	r1 = 10000
// 	r2 = 20000
// 	return
// }

func main() {
	c := fool("abc", 555)
	fmt.Println("c =", c)

	ret1, ret2 := fool2("haha", 999) // := 海象初始化
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = fool3("foo3", 333) // 初始化后，不能再使用 海象
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}
