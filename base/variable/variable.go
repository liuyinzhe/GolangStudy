package main

import (
	"fmt"
)

// 声明 全局变量，方法一、二、三 都可以
var g_A_var int = 100
var g_B_var = 200

func main() {
	//方法一 声明一个变量，int默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a=%T\n", a)

	//方法二 声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b=%T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s,type of bb =%T\n", bb, bb)

	// 方法三  初始化的时候可以省去 数据类型， 通过值自动匹配变量的数据类型
	var c = 100
	fmt.Print("c = ", c)
	fmt.Printf("type of c=%T\n", c)

	var cc string = "abcd"
	fmt.Printf("cc = %s,type of cc =%T\n", cc, cc)

	// 方法四 (常用方法) 省去 var 关键词 直接自动匹配
	// 不适用于 全局变量，只能用于函数体内
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e =%T\n", e)

	f := "abcd"
	fmt.Println("f =", f)
	fmt.Printf("type of f =%T\n", f)

	g := 3.14
	fmt.Println("f =", g)
	fmt.Printf("type of g =%T\n", g)
	// ===
	fmt.Println("g_A_var =", g_A_var, "g_A_var =", g_B_var)

	// 声明多变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk =", kk, "ll =", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, "jj =", jj)
}
