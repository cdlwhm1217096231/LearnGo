package main

import (
	"fmt"
	"math/rand"
)

func main(){
	// const，用于声明常量
		// 常量的值不可以改变
	// var，用来声明变量
		// 想要使用变量首先需要进行声明！！！
	const lightSpeed = 299792
	var distance = 5600000

	fmt.Println(distance / lightSpeed, "seconds")

	distance = 4010000
	fmt.Println(distance / lightSpeed, "seconds")

	// 同时声明多个变量
	var distance1 = 5600000
	var speed = 100000
	fmt.Println(distance1, speed)

	var (
		distance2 = 5600000
		speed2 = 100000
	)
	fmt.Println(distance2, speed2)

	var distance3, speed3 = 5600000, 100000
	fmt.Println(distance3, speed3)

	const hoursPerDay, minutesPerHour = 24, 60

	// 赋值运算符
	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783

	// 自增运算符,GO中无前置++运算符
	var age = 41
	age = age + 1
	age += 1  // 建议：使用age++来替换age += 1
	age++
	// ++age 错误，go语言中无前置++运算符

	// 猜数
		// 使用rand包，可以生成伪随机数。例如：Intn可以返回一个指定范围的随机整数
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
	
}