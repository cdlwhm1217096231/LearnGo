package main

import "fmt"

func main(){
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")
	fmt.Println()
	fmt.Println("My weight on the surface of Mars is", 149.0 * 0.3783, "lbs, and I would be", 41 * 365.2425/687, "years old.")
	// 格式化打印
	fmt.Println("------------------------------------格式化打印--------------------------------")
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0 * 0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41 * 365 / 687)
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
	fmt.Println("-------------使用Printf对齐文本-----------")
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
	
}
// Print、Println函数可以传递若干个参数，参数之间可以用逗号隔开
// 参数可以是字符串、数字、数学表达式等
// 格式化打印：可以使用Printf来控制打印的输出结果，与Print和Println不同，Printf的第一个参数必须是字符串。
	// 这个字符串中包含了像%v这样的格式化动词，它的值由第二个参数的值所代替。
	// 如果指定了多个格式化动词，那么它们的值由后面的参数值按顺序进行替换。

// 使用Printf对齐文本
	// 在格式化动词中指定宽度，就可以对齐文本。例如：%4v就是向左填充足够4个宽度
		// 正数：向左填充空格
		// 负数：向右填充空格
	
