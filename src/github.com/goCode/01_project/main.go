package main

import "fmt"

func main() {
	fmt.Print("hello world")
}


// CurryCoder.com是个人的域名，StudyGolang是不同的业务组，01_project是项目文件夹
// go mod init

// 在项目文件夹下打开终端，输入go run main.go就可以执行程序了
// 在项目文件夹下打开终端，输入go build就可以编译程序，再运行对应的xxx.exe文件(即项目文件夹的名字)即可
// 在项目文件夹下打开终端，输入go build -o dana.exe就可以编译出指定名称的.exe可执行文件，再运行对应的xxx.exe文件(即包的名字)即可

/*  交叉编译设置  */
// SET CGO_ENABLED=0   禁用CGO
// SET GOOS=linux   目标平台是linux
// SET GOARCH=amd64  目标处理器架构amd64


// go install
	// 分两步：先运行go build，然后将生成的xxx.exe文件自动移动到bin文件夹下

