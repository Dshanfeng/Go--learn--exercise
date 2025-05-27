package main

import "fmt"

func main() {
	//声明数组  var 数组名 [数组长度]数据类型
	//数组长度也属于数组类型的一部分，长度不同数据类型相同的也不同一类型
	//var arr1 [3]int
	//for i := 0; i < len(arr1); i++ {
	//	fmt.Println("请输入第", i, "个数的值：")
	//	fmt.Scanln(&arr1[i])
	//}
	//for _, v := range arr1 {
	//	fmt.Printf("%d\t", v)
	//}
	//fmt.Println()

	fmt.Println("---------")

	//初始化数组
	var arr2 [3]int = [3]int{2, 3, 4}
	fmt.Println(arr2)
	var arr3 = [3]int{8, 7, 6}
	fmt.Println(arr3)
	var arr4 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4) //根据你输入的值确定数组长度
	var arr5 = [...]int{3: 4, 0: 1}
	fmt.Println(arr5) //指定索引赋值，未指定的就是0

	fmt.Println("-----------")

	//Go语言中数组属于值类型而非引用类型
	var arr6 = [...]int{1, 2, 3, 4, 5}
	//arr6存储了整个数组，arr6的地址指向了整个数组的顶部
	fmt.Printf("arr6的类型：%T,arr6=%v,&arr6=%p\n", arr6, arr6, &arr6)
	for i := 0; i < len(arr6); i++ {
		//每个空间占用字节数有数据类型决定，int是64位8字节，每个地址与上一个差8
		fmt.Printf("arr6[%d]=%d, &arr6[%d]=%p\n", i, arr6[i], i, &arr6[i])
	}

	fmt.Println("---------")

	//内存空间在栈中，函数参数使用的是值传递
	//想要直接更改数组内容，请传递地址
	set(&arr6)
	fmt.Println(arr6)

	fmt.Println("---------")

	//二维数组
	var arr7 = [2][3]int{{1, 2, 3}, {7, 8, 9}}
	fmt.Printf("arr7的类型：%T, arr7=%v, &arr7=%p\n", arr7, arr7, &arr7)
	for key, value := range arr7 {
		fmt.Printf("key=%d,value=%d\n", key, value)
		for i, v := range value {
			fmt.Printf("arr7[%d][%d]=%d\t", key, i, v)
		}
		fmt.Println()
	}

}

func set(arr *[5]int) {
	(*arr)[0] = 100
}
