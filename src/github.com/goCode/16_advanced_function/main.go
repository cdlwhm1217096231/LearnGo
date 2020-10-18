package main


import (
	"fmt"
	"math/rand"
	"time"
)

type floor float64

func fakeSensor() floor {
	return floor(rand.Intn(151) + 150)
}

func realSensor() floor {
	return 0
}


func measureTemperature(samples int, sensor func() floor){
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v°K\n", k)
		time.Sleep(time.Second)
	}
}

var f = func() {
	fmt.Println("Dress up for the degree.")
}


// 因为函数字面值需要保留外部作用域的变量引用，所以函数字面值都是闭包的。
type dana float64
type curry func() dana

func realAnd() dana {
	return 0
}

func calibrate(c curry, offset dana) curry {
	return func() dana {
		return c() + offset
	}
}
func main(){
	// 高阶函数或一等函数
		// 1.在go语言中，函数是头等的，它可以用在整数、字符串或其他类型能用的地方：
			// a.将函数赋值给变量
			// b.函数作为参数传递给函数
			// c.将函数作为函数的返回类型


	// a.将函数赋值给变量

	// 变量sensor就是一个函数，而不是函数执行的结果

	// 无论sensor的值是fakeSensor还是realSensor，都可以通过sensor()来调用

	// sensor这个变量的类型是函数，该函数没有参数，返回一个floor类型的值
	sensor := fakeSensor  // 等价于 var sensor func() floor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	// b.将函数传递给其它函数
	measureTemperature(3, fakeSensor)

	// 声明函数类型
		// 为函数声明类型有助于精简和明确调用者的代码：type sensor func() floor
	
	// 闭包和匿名函数
		// 1.匿名函数：没有名字的函数，在go语言中也称为函数字面值
	f()

	ff := func(message string) {
		fmt.Println(message)
	}
	ff("Go to the party.")

	// 匿名函数
	func(){
		fmt.Println("Functions anonymous")
	}()
	
	// 闭包：由于匿名函数封闭并包围作用域中的变量而得名的。
	curry := calibrate(realAnd, 5)
	fmt.Println(curry())
}