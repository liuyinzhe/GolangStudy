package main

import "fmt"

//const 定义枚举类型
const (
	// 添加关键字  iota ,每行 iota都会累加1，第一行iota的默认值时0
	BEIJING  = 10 * iota // iota =0
	SHANGHAI             // iota =1 //10
	SHENZHEN             // iota =2 //20
)

const ( // iota 只能配合const
	a, b = iota + 1, iota + 2 // iota = 0  a=1 b =2
	c, d                      // iota =1 c=1,d=3
	e, f                      // iota =2 e=3,d=4
	g, h = iota * 2, iota * 3 // iota =3 ,g =6,h=9
	i, k                      // iota =4 ,i =8,k=12
)

func main() {
	// 常量(只读属性)
	const length int = 10
	fmt.Println("length = ", length)
	// const ()
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

}
