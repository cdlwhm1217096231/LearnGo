package main

import "fmt"


// 函数声明
	// go在标准库文档中列出了标准库每个包中声明的函数

	// 使用func关键字声明函数
	// 函数声明的一般格式：func 函数名 (变量名 类型名) 返回值类型

// 函数声明
	// 在go语言中，大写字母开头的函数、变量或其他标识符都会被导出，对其它包也可用
	// 小写字母开头的就不行！

// 函数声明之多个参数
	// func Unix(sec int64, nsec int64) Time

// 函数声明时，如果多个参数形参类型相同，那么该类型只写一次就可以了
	// func Unix(sec, nsec int64) Time

// 函数声明之返回多个值
	// go的函数可以返回多个值，如: countdown, err := strconv.Atoi("10")

// 函数的多个返回值需要用括号括起来，每个返回值名字在前、类型在后。声明函数时可以把名字去掉，只保留类型。
	// func Atoi(s string) (i int, err error) 等价于func Atoi(s string) (int, error)

// 可变函数参数
	// Println()是一个特殊的函数，它可以接收1个、2个甚至是多个参数，参数类型还可以不同。
	// fmt.Println("Hello World")
	// fmt.Println(185, "seconds")
// Println()的声明如下：
	// func Println(a...interface{}) (n int, err error)

// ...表示函数的参数数量是可变的
// 参数a的类型是interface{}，是一个空接口


// 函数按值进行传递
// 同一个包中声明的函数在调用彼此时不需要加上包名。
func sayHello(name string) string {
	return "Hello " + name
}


func main(){
	name := "CurryCoder"
	res := sayHello(name)
	fmt.Println(res)
}