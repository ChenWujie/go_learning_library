package main

import "fmt"

const (
	id  = "name" //只有变量不可以定义后不使用
	key = 0      //无类型变量，可以赋值给int，float，byte
)

func main() {
	fmt.Println("hello world!")

	var flag bool
	var isAwesome = true
	fmt.Println(flag, isAwesome) // false true

	var q = 10
	var w int
	fmt.Println(q, w) // 10 0

	var a = "hello"
	var b = a[0]
	var c string = "hello"
	fmt.Println(a, b, a == c) // hello 104 true

	var character rune
	var character2 = 'a'               //字符（rune）是int32的别名，byte是uint8的别名
	fmt.Println(character, character2) // 0 97

	var ( //声明多个变量时，可用var括号的形式
		x            = 10
		y    float64 = 20.2
		z            = float64(x) + y
		d            = x + int(y)
		e, f         = "hello", "world!"
	)
	fmt.Println(z, d) // 30.2 30
	fmt.Println(e, f) // hello world!

	f, g := ":=运算符还有用var关键字无法完成的功能", "允许对已经赋值的变量再次赋值，只要左边有一个新变量就行"
	fmt.Println(f, g)

}
