package main

import "fmt"

// 闭包Closure,将函数与其引用的量捆绑形成组成
// 被捆绑的量作用域为整个周期一直存活
// 全局变量也可以满足这样的结果为什么会有闭包，闭包又是怎么实现的呢？
var sum = 0

func main() {
	fmt.Println("全局变量的效果：")
	fmt.Println(getSum(1))
	fmt.Println(getSum(2))
	fmt.Println(getSum(3))
	fmt.Println(getSum(4))
	fmt.Println(getSum(5))

	fmt.Println("闭包的效果：")
	f := getSumClosure()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))

	//两者的效果是一致的，闭包的优势在于，其有私有化属性的效果，我们不能改变闭包捆绑的量(参数和局部变量)
	//全局变量是可以更改的，这点是全局变量没有的优势
	fmt.Println("---------")
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Printf("double的类型：%T, double(3)的类型：%T, double(3) = %d", double, double(3), double(3))
	fmt.Printf("triple的类型：%T,triple(3)的类型：%T，triple（3）= %d", triple, triple(3), triple(3))

	//多参数函数，通过闭包转换成了单参数函数
	//闭包可以将外层参数固定绑定以减少参数，例如：字符串转换整数时，进制就可以作为其参数
}

func getSum(mun int) int {
	sum += mun
	return sum
}

func getSumClosure() func(num int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func multiplier(factor int) func(num int) int {
	return func(num int) int {
		return num * factor
	}
}
