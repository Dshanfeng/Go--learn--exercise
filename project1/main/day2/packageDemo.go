package main

import (
	"fmt"
	"project1/main/day2/testPackage1"
	"project1/main/day2/testPackage2"
	//可以给包取别名，取过别名的包就不能在使用原包名了
	p3 "project1/main/day2/testPackage2/testPackage3"
)

func main() {
	//包：在Go语言中不同的函数定义都在不同的包中
	//包名可以与所处文件夹名不同，这点与Java有所不同
	//Go是不支持重载的，我们调用别的包内的函数或变量都是通过包名调用的，这意味着同一个包下不允许存在同名的东西
	//所以同名函数我们可以定义在不同的包中，来实现所谓的重载效果

	//在程序层面，所有使用相同package包的源文件组成的代码模块就是包
	//在文件层面包就是一个文件夹

	//不需要在意包，当成与目录同名一样正常使用就行了，遇到问题了在总结在处理

	fmt.Println("Go语言包实验开始，相关目录testPackage1、testPackage2和testPackage3")

	//与目录同名的包
	fmt.Println(testPackage1.Add(3, 4))

	//与目录不同名的包
	fmt.Println(package2.Multiply(3, 4))
	fmt.Println(package2.Subtract(3, 4))

	fmt.Println(p3.Remainder(3, 4))
}
