package main

import (
	"fmt"
	"strconv"
)

func main() {
	//类型转换
	//go没有隐式类型转换，只有强制类型转换
	var n1 int8 = 2
	var f1 float32 = float32(n1)
	fmt.Println(f1)
	fmt.Printf("f1的类型%T,f1 = %f\n", f1, f1) //他们打印的格式取决与占位符

	var f2 float32 = 3.14
	var n2 int8 = int8(f2)
	fmt.Printf("n2的类型%T，n2 = %d\n", n2, n2)

	// n3 := 22000
	// var b bool = bool(n3)r
	// fmt.Println("b的类型%T,b = %t",b,b)
	//错误的类型转换，其他类型是转不了布尔类型的

	//基本数据类型转字符串有两种方法
	//第一种，Sprintf返回值
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("%q\n", s1) //%q

	//第二种，strconv包的方法，由于方法较麻烦不建议使用
	var s2 string = strconv.FormatInt(int64(n2), 10)
	fmt.Printf("%q", s2)
	var s3 string = strconv.FormatFloat(float64(f1), 'f', 3, 32)
	fmt.Printf("%q", s3)
}
