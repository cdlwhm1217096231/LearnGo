package main


import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	// 类型不能混合使用

	// 1.连接两个字符串，使用+运算符
	countdown := "Launch in T minus " + "10 seconds."
	fmt.Println("countdown =",countdown)

	// 2.如果想连接字符串和数值，是会报错
	// countup := "hello " + 111 + " world."   报错！！！
	// fmt.Println(countup)

	// 3.整型和浮点型不能混着用
	// age := 18
	// marsDays := 687
	// earthDays := 365.982
	// fmt.Println("I am", age * earthDays / marsDays, "years old on Mars.") 报错

	// 4.数值类型之间的转换
	age := 18
	marsAge := float64(age)
	fmt.Println(marsAge)

	// 5.浮点类型转换为整数类型
		// 可以从浮点类型转化为整数类型，小数点后面的部分会被截断，而不是四舍五入
	earthDays := 692.4343
	fmt.Println(int(earthDays))

	// 6.无符号和有符号整数类型之间的转换

	// 7.不同大小的整数类型之间也需要转换

	// 8.类型转换时需谨慎，会发生环绕行为
	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h)

	// 可以通过math包提供的max、min常量，来判断是否超过最大和最小值
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("handle out of range value")
	}

	// 9.字符串转换
	// 把rune、byte转换为string
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 22
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	// 把数值转换为string，它的值必须能转换为code point
	fmt.Println(string(65))
	fmt.Println(string(4545453434342))

	// strconv包的Itoa函数就可以将数值转换为string
	countdown1 := 10
	str := "Launch in T minus " + strconv.Itoa(countdown1) + " seconds."
	fmt.Println(str)

	// Itoa是Integer to ASCII的意思
	// Unicode是ASCII的超集，它们前128个code points是一样的

	// 另外一种把数值转换为string的方式是使用Sprintf函数，和Printf类似，但是会返回一个string
	cout := 9
	str1 := fmt.Sprintf("Launch in T minus %v seconds.", cout)

	fmt.Println(str1)
	fmt.Printf("Type of %s is %T\n", str1, str1)

	// strconv包中还有一个Atoi函数
	// 由于字符串中可能包含任意字符，或者要转换的数字字符串太大，所以Atoi函数可能会发生错误。

	// 如果err的值为nil,则代表没发生错误。
	cin, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("no errors")
	}
	fmt.Println(cin)

	// go是静态类型语言，一旦某个变量被声明，那么它的类型就无法再改变了。

	// 布尔类型
		// Print家族函数中，会把bool类型的值打印成true/false的文本
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch: ", yesNo)

	// 注意：如果想使用string(false),int(false),bool(1),bool("yes")等类似的方式进行转换，那么go编译器会报错
	// 某些语言中，会将1和0当成true和false，但是go语言中不行
	
}