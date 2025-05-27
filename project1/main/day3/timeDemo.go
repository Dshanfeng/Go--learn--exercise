package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("now的类型：%T,%v\n", now, now)
	fmt.Println("-------")

	//这种标准格式不适合观看，所以我们一般格式化处理，返回字符串
	str := now.Format("2006-01-02 15:04:05")
	fmt.Printf("str的类型：%T,%s\n", str, str)
	//这些日期便是Go语言的生日，每个数字都对应一个单位，这些数字便像是占位符，你可以任意组合
	fmt.Println(now.Format("2006年01月02日  15点04分"))
	fmt.Println("---------")

	//通过方法可以获取你想要的部分，结合Sprintf也可以做到格式化字符串
	str = fmt.Sprintf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(str)
	fmt.Println("--------")

	//展示一下这些函数的返回值
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	//月份这里有个特点会打印英文
	fmt.Println(int(now.Month())) //强转成数字即可
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//内置函数，new，用于动态分配内存，返回的是指针类型
	//new 出来的值都是在内存在堆中，Go有内存管理机制，不用手动释放内存
	nPoint := new(int)
	*nPoint = 100
	fmt.Printf("nPoint的类型：%T, nPoint的值（指向地址）：%v，nPoint的地址：%v，nPoint指向的值：%v\n",
		nPoint, nPoint, &nPoint, *nPoint)

}
