package main

import (
	"fmt"
	"project1/main/day3/initTest"
)

// 匿名函数，只使用一次的函数，定义加调用一起使用
var test = func() string {
	fmt.Println("全局变量test初始化")
	return "test"
}()

// 匿名函数还有一种使用方法，将其赋值给变量，同时这种方法可以多次使用了
var anonyFunc = func(num1 int, num2 int) int {
	return num1 * num2
}

// 初始化函数，在main函数之前调用，作用为变量初始化，优先级低于直接初始化
func init() {
	fmt.Println("调用了initDemo.go中的init函数")
}

func main() {
	fmt.Println("调用了initDemo.go中的main函数")
	fmt.Println(test)
	fmt.Println(initTest.InitDemoTest)

	fmt.Println("-----------")
	fmt.Println("匿名函数的测试:")
	fmt.Printf("anonyFunc的类型：%T, anonyFunc(2,3) = %d", anonyFunc, anonyFunc(2, 3))
}

//在本例中可以发现，整个执行过程是这个样子的，
//1、导包，将引用包的全局变量初始化 ，调用其init函数
//2、调用本文件的init函数
//3、调用main函数
