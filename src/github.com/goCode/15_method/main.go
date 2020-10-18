package main

import "fmt"


// 可以将方法与同包中声明的任何类型相关联，但不可以是int、float64等预声明的类型进行关联。
type cell float64
type floor float64

func convertT(f floor) cell {
	return cell(f - 273.15)
}

// 将cell()方法与floor类型相关联
func (f floor) cell() cell {
	return cell(f - 273.15)
}

// 上述cell()方法中，虽然没有参数，但它前面却有一个类型参数的接收者f

// 每个方法可以有多个参数，但只能有一个接收者

// 在方法体中，接收者的行为和其它参数一样


// 方法的声明
	// func (f floor) cell() cell
		// f floor：接收者,可以将方法与同包中声明的任何类型floor相关联，但不可以是int、float64等预声明的类型进行关联。
		// cell()：方法名称
		// cell：返回类型


func main(){
	// 关键字type可以用来声明新的类型
		// 例如：type celsius float64
	type celsius float64
	var temperature celsius = 20
	// 虽然celsius是一种全新的类型，但是由于它和float64具有相同的行为和表示，所以赋值操作能顺利执行。
	temperature += 20;
	fmt.Println(temperature)

	// 为什么需要声明新类型？因为这样能极大的提高代码的可读性和可靠性

	// 不同类型是无法混用的
	// var warmUp float64 = 10
	// temperature += warmUp  报错！
	// fmt.Println(temperature)

	// 通过方法添加行为
		// 1.在C++ Java中，方法属于类
		// 2.在go语言中，它提供了方法，但是没有提供类和对象
		// 3.go语言比其他语言的方法要灵活

	// 方法调用：变量.方法()
	var f floor = 294.0
	var c cell

	c = convertT(f)
	fmt.Println(c)
	
	c = f.cell()
	fmt.Println(c)
}