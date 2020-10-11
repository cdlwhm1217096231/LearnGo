package main

import (
	"fmt"
	"math"
	"time"
)

func main(){
	/*  go语言中的整数类型  */
	// 1. go语言提供了10种整数类型
		// a.不可以存小数部分
		// b.范围有限
		// c.通常根据数值范围来选择整数类型
	// 2. 5种整数类型是有符号的
		// a.能表示正数、0、负数
	// 3. 5种整数类型是无符号的
		// a.能表示正数、0
	
	// 最常用的整数类型int,下面三个语句是等价的
	var year int = 2020
	// year := 2020
	// var year = 2020
	fmt.Println("year =", year)

	// 无符号类型整数是uint
	var haha uint = 2
	fmt.Println("haha =", haha)

	/* 8种整数类型 */
	// 整数类型包括有符号和无符号，实际上一共有8种类型。
		// a.它们的取值范围各不相同
		// b.与架构无关
	// int8   -128-127
	// uint8  0-255
	// int16  -32768-32767
	// uint16 0-65535
	// int32  uint32
	// int64  uint64

	/* int 和 uint */
	// int和uint是针对目标设备优化的类型：
		// a.在树莓派2、比较老的移动设备上，int和uint32都是32位
		// b.在比较新的计算机上，int和uint都是64位的

	// 虽然在某些设备上int可以看成是int32，在某些设备上int也可以看成是int64，但是它们实际是3种不同类型

	// int并不是其它类型的别名

	/* 打印数据类型 */
	// 在Printf中使用%T就可以打印数据的类型
	fmt.Printf("Type %T for %v\n", year, year)

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	// uint8可以用来表示8位的颜色(0-255)
	var r, g, b uint8 = 0, 141, 213
	fmt.Println(r, g, b)

	// 十六进制表示法
		// go语言中，在数前面加上0x前缀，就可以用十六进制的形式表示数值
	r, g, b = 0x00, 0x8d, 0xd5
	// 打印十六进制数，使用%x格式化动词
	fmt.Printf("%x %x %x\n", r, g, b)
	// 也可以指定最小宽度和填充
	fmt.Printf("#%02x%02x%02x\n", r, g, b)

	/* 整数环绕 */
	// 所有的整数类型都有一个取值范围，超过这个范围就会发生环绕。
	var red uint8 = 255
	red++
	fmt.Println(red) // 0

	var num int8 = 127
	num++
	fmt.Println(num)  // -128

	/*  打印每个bit */
	var green uint8 = 3
	// %b格式化动词
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	/* 整数类型的最大和最小值 */
	// Math包中，为与架构无关的整数类型，定义了最大和最小值常量
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MinInt64)

	// 而int和uint，可能是32位或64位的
}