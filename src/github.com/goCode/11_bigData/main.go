package main

import (
	"fmt"
	"math/big"
)

func main(){
	// 数太大了怎么办？
		// 1.浮点类型可以存储非常大的数值，但精度不高
		// 2.整型很精确，但取值范围有限
		// 3.如果需要很大的数，而且要求精度很高，该怎么办？：
			// a.int64可以容纳很大的数，如果还不行，则：
				// b.uint64可以容纳更大的正数，如果还不行，则：
					// c.也可以凑合用浮点类型，但还有另一种方法：“使用big包”
	var dis int64 = 42.3e12
	const speed = 299792
	const secondsPerDay = 86400
	fmt.Println("distance =", dis)
	days := dis / speed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.")
	
	// 如果没有为指数形式的数值指定类型，则go语言中会将其视为float64类型
	var distance = 24e18
	fmt.Println(distance)


	/* big  */
	// 对于较大的整数(超过10^18)：big.Int
	// 对于任意精度的浮点类型：bigFloat
	// 对于分数:big.Rat
	speeds := big.NewInt(299792)
	secondsPerDays := big.NewInt(87634)
	fmt.Println(speeds, secondsPerDays)
	// 如果NewInt()中传入的数超过int64范围，则使用下面的方法
	demo := new(big.Int)
	demo.SetString("24000000000000000", 10)
	fmt.Println(demo)

	// 一旦使用了big.Int,那么等式中其它的部分也必须使用big.Int
	// NewInt()函数可以把int64转换成big.Int类型

	/* 较大数值的常量 */
	// 在go语言中，可以指明常量类型，也可以不指明常量类型
	// 对于变量，go使用类型推断
	// 在go语言中，常量是可以无类型的
	// const distanceff int64 = 12343000000000000000  报错
	const distanceff = 12343000000000000000  
	// fmt.Println(distanceff)

	// 常量使用const关键字来声明，程序中的每个字面值都是常量。这意味着比较大的数值可以直接使用

	// 针对字面值和常量的计算是在编译阶段完成的
	


}