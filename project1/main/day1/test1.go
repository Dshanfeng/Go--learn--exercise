package main

import (
	"fmt"
)

func main() {
	var n1 int8 = 1
	var n2 int16 = 128
	var n3 int32 = 33000
	var n4 int64 = 2200000000
	fmt.Printf(" n1 = %d \n n2 = %d \n n3 = %d \n n4 = %v", n1, n2, n3, n4)
	//GO语言特性，凡是定义的变量必须使用，不使用就会报错
	//导包也是如此，必须使用

	var n5 = 12 //根据赋值自动判断变量类型，整数类型的默认类型int,8字节64位
	fmt.Printf("整数类型的默认类型是：%T,n5 = %d\n", n5, n5)

	n6 := 22 //更简单的写法，省略了var声明
	fmt.Printf("n6 = %d\n", n6)

	n7 := 22000000000
	fmt.Println(n7)

	var n8 uint8 = 3 //无符号整型
	fmt.Println(n8)

	f1 := 3.14 //float64
	fmt.Printf("浮点数类型的默认类型是：%T,f1 = %f\n", f1, f1)

	var f2 float32 = 3.40312345
	fmt.Println(f2)

	var b bool = false
	fmt.Println(b)

	var c1 byte = 'a' //字符类型，Go语言没有专门的字符类型，使用整数类型byte代替，打印时有些麻烦
	fmt.Println(c1)
	fmt.Printf("c1 = %c\n", c1) //第一种：格式化输出，占位符%c

	c2 := 'b'
	fmt.Printf("字符类型的默认类型是：%T\n", c2)

	c3 := `c` //第二种：反引号
	fmt.Println(c3)

	s1 := "Hello World!"
	fmt.Println(s1)

	z := &s1 //指针，复杂数据类型
	fmt.Printf("z的数据类型：%T，z = %p", z, z)

}
