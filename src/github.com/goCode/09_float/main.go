package main

import (
	"fmt"
	"math"
)

func main(){
	// 声明浮点类型变量
	// 以下三种声明方式均是等价的
	days := 3.1415
	// var days = 3.1415
	// var days float64 = 3.1415
	// 只要数字含有小数部分，那么它的类型就是float64
	fmt.Println("days =", days)
	// 如果你使用一个整数来初始化某个变量，那么你必须指定它的类型为float64,否则它就是一个整数类型
	var answer float64 = 42
	fmt.Printf("answer = %f\n", answer)

	year := 2020
	fmt.Println("year =", year)

	/* 单精度浮点数类型 */
	// Go语言中有两种浮点数类型：
		// 1.默认是float64
			// a.64位的浮点类型
			// b.占用8个字节内存
			// c.某些编程语言把这种类型称为double
		// 2.float32
			// a.占用4个字节内存
			// b.精度比float64低
			// c.有时候也叫单精度类型
	
	// 想使用单精度类型，必须在声明变量时指定该类型
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println("pi64 =", pi64)
	fmt.Println("pi32 =", pi32)

	/* 单精度的使用场景 */
		// 1.当处理大量数据时，例如3D游戏中的数千个顶点，使用float32牺牲精度来节省内存是有意义的
		// 2.math包中的函数操作都是float64类型，所以应该首选使用float64，除非你有足够的理由不去使用它
	
	
	/* 零值 */
		// go语言中每个类型都有一个默认值，称为零值
		// 当你声明变量却不对其进行初始化时，它的值就是零值
	var price float64
	fmt.Println("price =",price)

	/* 显示浮点类型 */
		// 1.使用Print或Println打印浮点类型时，默认的行为是尽可能多显示几位小数
		// 2.如果你不想这样，那么应该使用Printf函数，结合%f格式化动词来指定显示小数的位数
	third := 1.0 / 3
	fmt.Println("third =",third)
	fmt.Printf("thirdv = %v\n", third)
	fmt.Printf("thirdf = %f\n", third)
	fmt.Printf("third3f = %.3f\n", third)  // %.3f小数点后保留3位
	fmt.Printf("third42f = %4.2f\n", third) // %4.2f包含小数点共4位，其中小数点后保留2位

	second := 11 / 3
	fmt.Println("second =",second)  // 3

	/* %f格式化动词 */
	// 它有两部分组成：宽度 + 精度
		// 宽度：会显示出的最少字符个数(包含小数点和小数)
			// a.如果宽度大于数字的个数，那么左边会填充空格
			// b.如果没有指定宽度，那么就按实际的位数进行显示
		
		// 精度：小数点后面显示的位数

	// 如果想使用0代替空格作为填充
	fmt.Printf("%05.2f\n", third)

	first := 21.0 / 5.0
	fmt.Println(first)

	// 如何比较两个浮点数
	payment := 0.1
	payment += 0.2
	fmt.Println(payment == 0.3)

	// 正确的做法
	fmt.Println(math.Abs(payment - 0.3) < 0.0001)
}