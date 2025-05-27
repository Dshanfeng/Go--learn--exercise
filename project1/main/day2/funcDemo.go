package main

import (
	"fmt"
)

func main() {
	num1 := add(1, 2)
	fmt.Println(num1)
	fmt.Println("---------")

	sum, avg := average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("和：%d,平均值：%f\n", sum, avg)
	fmt.Println("---------")

	sum1, avg1 := average1(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("和：%d,平均值：%f\n", sum1, avg1)
	fmt.Println("---------")

	var n1, n2 int8 = 3, 7
	fmt.Printf("交换前：n1=%d,n2=%d\n", n1, n2)
	exchageNums(&n1, &n2)
	fmt.Printf("交换后：n1=%d,n2=%d\n", n1, n2)
	fmt.Println("---------")

	//在Go语言中函数也是一种数据类型m,一个函数可以赋值给一个变量
	func1 := add
	fmt.Printf("func1的数据类型是：%T\n", func1)
	fmt.Println("-------")

	//函数类型就是 func（参数类型列表）（返回值类型列表）
	var func2 func(*int8, *int8) = exchageNums
	//这个函数类型的变量和函数具有一样的功能，当函数名很长时我们就可以使用它
	var n3, n4 int8 = 5, 4
	fmt.Printf("交换前：n1=%d,n2=%d\n", n3, n4)
	func2(&n3, &n4)
	fmt.Printf("交换后：n1=%d,n2=%d\n", n3, n4)
	fmt.Println("---------")

	var n5, n6, n7, n8 int8 = 1, 2, 3, 4
	fmt.Printf("交换前：n5=%d,n6=%d,n7=%d,n8=%d\n", n5, n6, n7, n8)
	exchageTest(func2, &n5, &n6, &n7, &n8)
	fmt.Printf("交换后：n5=%d,n6=%d,n7=%d,n8=%d\n", n5, n6, n7, n8)
	fmt.Println("---------")

	//自定义数据类型--起别名
	type myString string
	var str myString = "你好！"
	fmt.Printf("str的数据类型：%T , str = %s\n", str, str)
	//虽然是取别名，但不能将其直接赋值给原类型，需要强转
	var str2 string = string(str)
	fmt.Printf("str2的数据类型：%T , str2 = %s\n", str2, str2)

}

func add(a int, b int) int {
	return a + b
}

// Go语言函数可以实现可变参数，多返回值（多返回值要用括号扩起，返回时必须一一对应）
func average(nums ...int) (int, float64) {
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total, float64(total) / float64(len(nums))
}

// Go语言还可以指定返回值变量，就不用一一对应了
// Go语言不支持重载
func average1(nums ...int) (sum int, avg float64) {
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	avg = float64(sum) / float64(len(nums))
	return
}

// Go语言基本数据类型是值传递，操作的是其副本，但Go语言存在指针操作，可以实现地址传递
// 无返回值就不用写返回值
func exchageNums(num1 *int8, num2 *int8) {
	var temp int8 = *num1
	*num1 = *num2
	*num2 = temp
}

// 由于函数也有类型因此函数也可以作为参数
func exchageTest(testFunc func(*int8, *int8), nums ...*int8) {
	for i := 0; i < len(nums)-1; i += 2 {
		testFunc(nums[i], nums[i+1])
	}
}
