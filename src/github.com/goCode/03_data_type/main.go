package main 

import (
	"fmt"
	"math"
	"Strings"
)

func main(){
	// uint8就是byte类型
	// int16对应C语言中的short类型
	// int64对应C语言中的long类型
	// int和uint自动匹配特定平台整型长度的类型
	// len()获取对象的长度
	var a int = 10
	// 十进制
	fmt.Printf("%d \n", a)
	// 二进制
	fmt.Printf("%b \n", a)
	// 八进制
	var b int = 077
	fmt.Printf("%o\n", b)
	// 十六进制
	var c int = 0xff
	fmt.Printf("%x\n", c)
	fmt.Printf("%X\n", c) 
	// 变量a的内存地址(十六进制表示)
	fmt.Printf("%p\n", &a)

	/*   浮点类型   */
	// float32和float64
	// float32的浮点数最大范围约为math.MaxFloat32
	// float64的浮点数最大范围约为math.MaxFloat64
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)

	/*  布尔值  */
	// 布尔类型变量的默认值是false
	// Go语言中不允许将整型强制转换为布尔类型
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	
	/* 字符串  */
	// go语言中的字符串内部实现使用utf-8编码
	// 多行字符串使用反引号``
	fmt.Println("\"c:\\go\"")
	var s1 = `
		这是
		多行
		文本
	`
	fmt.Println(s1)

	s2 := "CurryCoder"
	// 计算长度
	fmt.Println(len(s2))
	// 拼接字符串
	s3 := " Hello World!"
	fmt.Println(s2 + s3)
	
	s4 := fmt.Sprintf("%s-----%s", s2, s3)
	fmt.Println(s4)

	// 字符串分割
	ret := strings.Split(s2, "y")
	fmt.Println(ret)

	// 判断是否包含
	ret2 := strings.Contains(s2, "y")
	fmt.Println(ret2)

	// 判断前缀和后缀
	ret3 := strings.HasPrefix(s2, "Curry")
	fmt.Println(ret3)
	ret4 := strings.HasSuffix(s2, "Coder")
	fmt.Println(ret4)

	// 求子串的位置
	s5 := "applepen"
	fmt.Println(strings.Index(s5, "p"))
	fmt.Println(strings.LastIndex(s5, "p"))

	// join
	a1 := []string{"Python", "PHP", "C++", "Java", "JavaScript", "Golang"}
	fmt.Println(strings.Join(a1, "-"))

	// byte类型代表了一个ASCII编码的字符和rune类型代表一个utf-8字符
	ss1 := "Golang"
	cc1 := 'G' // ASCII码下占1个字节
	fmt.Println(ss1, cc1)

	ss2 := "中国"
	cc2 := '中'  // utf-8编码下，一个中文汉字占3个字节
	fmt.Println(ss2, cc2)

	ss3 := "hello中国"
	fmt.Println(len(ss3))

	// 遍历字符串
	for i := 0; i < len(ss3); i++{
		fmt.Printf("%c\n", ss3[i])
	}
	fmt.Println()
	// for range循环是按照rune类型去遍历的
	for _, v := range ss3{
		fmt.Printf("%c\n", v)
	}
	fmt.Println()
	/* 修改字符串 */
	// 需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	ss4 := "big"
	// 强制类型转换
	byteS1 := []byte(ss4)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))
	fmt.Println()
	ss5 := "白鹿村"
	runeS1 := []rune(ss5)
	runeS1[2] = '原'
	/*  类型转换  */
	// Go语言中只有强制类型转换，没有隐式类型转换：T(表达式)
	fmt.Println(string(runeS1))
}
