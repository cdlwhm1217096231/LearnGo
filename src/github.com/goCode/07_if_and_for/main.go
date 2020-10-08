package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	// 布尔类型
		// true和false是go语言中两个已经声明好的常量
		// go语言只有true是真的，只有false是假的。和其他语言如js不同，js会将""这种字符串当作false，其他的字符串当作true
	
	// strings.Contains
		// 来自strings包的Contains()函数可以判断某个字符串是否包含另外一个字符串
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	// 比较运算符
		// 如果我们比较两个值，得到的结果也是true或false
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor) 

	// 使用if来做分支
	var commands = "go east"

	if commands == "go east"{
		fmt.Println("You head further up the mountain.")
	} else if commands == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	// 逻辑运算符
		// ||表示或、&&表示与，它们通常用来检测多个条件
	fmt.Println("The year is 2020, should you leap?")

	var year = 2020
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}

	// 取反逻辑运算符!
	var haveTorch = true
	var litTorch = false
	
	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}

	// 使用switch做分支
		// switch语句也可以对数字进行匹配
		// 还有一个fallthrough关键字，它用来执行下一个case的body部分，这一点与C++等语言不同！
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var commandss = "go inside"

	switch commandss {
	case "go east":
		fmt.Println("You head further up the moutain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}

	// fallthrough关键字
	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

	// 使用循环做重复
		// for关键字可以让你的代码重复执行
		// for后面没有跟条件，就是无限循环
		// 可以使用break来跳出循环
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second * 2)
		count--
	}
	fmt.Println("Liftoff!")
}	
