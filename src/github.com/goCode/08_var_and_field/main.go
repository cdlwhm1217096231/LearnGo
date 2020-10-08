package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main(){
	// 变量的作用域
		// 当变量被声明以后，它就进入了作用域(变量就变得可见了)
			// 只有变量在作用域内，你就可以访问它
			// 否则，访问它会报错
		// 变量声明的位置，决定了它的作用域

	// 作用域的好处
		// 可以在不同的作用域内使用相同名称的变量名

	// 在go语言中，作用域的范围就是{}之间的部分
	var count = 10
	for count > 0 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count--
	}

	// 短声明
		// 在go语言中，可以使用var来声明变量var count = 10。但是，也可以使用短声明方式来声明变量count := 10
		// 两种声明方式的效果是一样的
		// 不仅短声明语句更短，而且可以在无法使用var的地方使用

	// package作用域
		// era变量是在main函数外声明的
			// 它拥有package作用域
			// 如果main package有多个函数，那么era对它们都可见
		// 短声明不可以用来声明package作用域的变量
	
	year := 2018

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
	

}