package main

import "fmt"

//defer关键字，作用延迟执行其后语句，到函数返回时才执行
//应用场景：在创建资源后，直接用defer释放资源，这样就不用考虑后续资源释放了
//       问题:defer是在函数返回时才执行，一定要考虑好资源的释放时机，像IO流这种要随用随建，不用时立即释放
//defer实现原理：defer有一个独立的栈存放其语句,在函数返回前执行这个栈中的语句
//           注：栈是先进后出的，效果：先写的语句后执行
//               语句相关的量也会一起拷贝进入栈中，所以后面对其的更改是影响不了defer的语句的

func main() {
	defer fmt.Println("一")
	num1 := 6
	num2 := 6
	defer fmt.Printf("num1 = %d  num2 = %d  sum = %d\n", num1, num2, getSum2(num1, num2))

	num1 = 8
	num2 = 2
	fmt.Printf("num1 = %d   num2 = %d  sum = %d\n", num1, num2, getSum2(num1, num2))

}

// 前面说过同一个包下不能有同名的函数，突然发现全局变量与函数也不允许同名，毕竟函数也是一种数据类型
func getSum2(n1 int, n2 int) int {
	return n1 + n2
}
