package main

import "fmt"

func main() {
	//运算符，没有三目运算符

	//算数运算符
	var n1, n2 int8 = 1, 2
	fmt.Printf("n1的类型%T, n1=%d\n", n1, n1)
	fmt.Printf("n2的类型%T, n2=%d\n", n2, n2)

	fmt.Printf("n1 + n2 = %d\n", n1+n2)
	fmt.Printf("n1 - n2 = %d\n", n1-n2)
	fmt.Printf("n2 * n1 = %d\n", n2*n1)
	fmt.Printf("n2 / n1 = %d\n", n2%n1)
	fmt.Printf("n2 %% n1 = %d\n", n2%n1)
	fmt.Printf("n1++ n2-- , n1 = %d , n2 = %d\n", n1, n2)

	//逻辑运算符与关系运算符
	fmt.Printf("%t&&%t  结果：%t\n", n1 == n2, n1 != n2, n1 == n2 && n1 != n2)
	fmt.Printf("%t||%t 结果：%t\n", n1 > n2, n1 <= n2, n1 >= n2 || n1 < n2)
	fmt.Println(!true)

	//赋值运算符
	n1 += 5
	fmt.Printf("n1 += 5 , n1 = %d\n", n1)
	n2 *= 2
	fmt.Printf("n2 *= 2 , n1 = %d\n", n1)

	//其他运算符
	var point1 *int8 = &n1
	fmt.Printf("point1的类型%T, point1=%d\n", point1, *point1)

}
