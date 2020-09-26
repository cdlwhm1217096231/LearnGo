package main

import "fmt"

// 定义一个函数
func foo()(string, int){
	return "CurryCoder123", 2000
}




func main(){
	// 变量声明格式：var 变量名 变量类型

	// 声明变量必须使用!!!
	var name string
	var age int
	var isOk bool


	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	// 批量声明
	var (
		a string // ""
		b int  // 0
		c bool  // false
	)

	// 变量赋值
	a = "CurryCoder"
	b = 23
	c = true
	fmt.Println(a, b, c)
	
	// 变量声明且初始化
	var x  string = "old boy"
	fmt.Println(x)

	fmt.Printf("%s 哈哈哈 %d ", x, b)

	// 类型推导(编译器根据变量初始值的类型指定给变量)
	var y = 28
	fmt.Println(y)

	var z = "Curry Coder"
	fmt.Println(z)

	// 短变量声明：在函数内部，可以使用更简洁的 := 方式声明并初始化变量
	CurryCoder := "库里粉丝"  // var CurryCoder string = "库里粉丝"
	fmt.Println(CurryCoder) 

	// 调用foo函数
	// 匿名变量_：_用于接收不需要的变量值
	aa, _ := foo()
	fmt.Println(aa)

	_, bb := foo()
	fmt.Println(bb)

	// 不能重复声明同名变量
	var n = "Coder"
	// var n = "Dana" 错误
	fmt.Println(n)

	/*    常量：不允许被修改     */
	const PI = 3.1415
	const e = 2.71

	// 批量声明常量
	const (
		aaa = 100
		bbb = 1000
		ccc  // 1000
		ddd  // 1000
	)
	fmt.Println(aaa, bbb, ccc, ddd)

	/* iota 枚举 */
	// 遇到const iota就初始化为0
	// const中每新增一行变量声明iota就递增1
	// const声明中如果不写，就和上一行一样
	const (
		abc = iota // 0
		aabbcc // iota
		aaabbbccc // iota
		aaaabbbbcccc // iota
	)
	fmt.Println(abc, aabbcc, aaabbbccc, aaaabbbbcccc)

	const (
		n1 = iota  // 0
		n2  // 1
		_  // 2
		n4  // 3
	)
	fmt.Println(n1, n2, n4)

	const (
		n11 = iota  // 0
		n22 = 100  // 100
		n33 = iota // 2
		n44 = iota // 3
	)
	fmt.Println(n11, n22, n33, n44)

	const (
		p, q = iota + 1, iota + 2 // 1 2
		w, v // 2 3 
		l, k // 3 4
	)
	fmt.Println(p, q, w, v, l , k)
}