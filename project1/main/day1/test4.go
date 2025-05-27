package main

import (
	"fmt"
)

func main() {
	//语句结构 ，顺序、分支、循环

	var num int
	fmt.Println("请输入一个数：")
	fmt.Scanln(&num)

	//分支结构

	if num%2 == 0 {
		fmt.Println("num是偶数")
	} else {
		fmt.Println("num是奇数")
	}

	if num > 0 && num <= 5 {
		fmt.Println("num大于0小于等于5")
	} else if num <= 10 {
		fmt.Println("num大于5小于等于10")
	} else if num <= 50 {
		fmt.Println("num大于10小于等于50")
	} else {
		fmt.Println("num大于50")
	}

	fmt.Println("---------")

	//if可以实现声明变量
	if f := 3.14; f > 3 {
		//f若是float可以与int类型比大小，但相反f是int类型不可以与float类型值比大小
		fmt.Println("f大于3")
	}

	fmt.Println("--------")

	//Go语言switch没有穿透特性不用使用break，
	//case后可接多个值用逗号分隔，
	//case后的值必须与表达式结果的类型一致
	switch num {
	case 0:
		fmt.Println("num = 0")
	case 1, 2, 3:
		fmt.Println("num 为 1，2，3其中之一")
	default:
		fmt.Println("num 不是 0，1，2，3中任何一个")
	}

	fmt.Println("--------")

	//Go语言的switch可以当作if使用，
	//关键字fallthrough可以穿透下一层
	switch {
	case num >= 0 && num <= 10:
		fmt.Println("num大于0小于等于10")
	case num > 10 && num <= 20:
		fmt.Println("num大于10小于等于20")
		fallthrough
	case num == 0:
		fmt.Println("不建议使用fallthrough")
	default:
		fmt.Println("<UNK>")
	}

	fmt.Println("--------")

	//循环结构，Go只有for一种循环结构
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}

	//for range 遍历数组、切片、字符串、map及通道
	str := "Hello World!"
	//按字节获取，中文不能使用
	for index, value := range str {
		fmt.Printf("索引：%d，值：%c\n", index, value)
	}


	//标签，break、continue都可以与标签搭配退出或重新循环指定块
label:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 4 && j == 4 {
				break label
			}
			fmt.Println(i, j)
		}
	}
}
