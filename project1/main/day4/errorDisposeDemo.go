package main

import (
	"errors"
	"fmt"
)

func main() {
	// 错误处理：增强代码的健壮性
	//不建议忽略错误"_"，忽略错误你连哪里有错误都不知道
	// defer 语句会将其后面跟随的语句进行延迟处理。
	// 在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
	// recover 处理 panic 出现的错误，捕获错误，使程序不会崩溃，正常执行
	fmt.Println(division(4, 0))
	fmt.Println("程序正常执行")

	fmt.Println("---------")

	//这种方法捕获的错误都属于Go定义的错误，都是英文的,我们还可以自定义错误
	//errors包的New函数可以创建一个error对象
	result, err := division3(4, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
		//若是想让程序发生错误时中断，panic()将错误信息输出在控制台,当没有错误使，其参数便是nil零值
		panic(err)
	}
	fmt.Println("程序正常执行")

}

func division(num1 int, num2 int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("程序出现错误：", err)
		}
	}()
	return num1 / num2
}

// 注意：想要程序能不顾错误正常执行 recover 必须搭配 defer 使用。
//func division2(x int, y int) int {
//	err := recover()
//	if err != nil {
//		fmt.Println(err)
//	}
//	return x / y
//}

func division3(num1 int, num2 int) (int, error) {
	//将错误作为返回值返回，我感觉这种处理方式像是java的抛出处理，交由函数调用者处理
	//这种处理方法使程序也不会崩溃，正常执行
	if num2 == 0 {
		return 0, errors.New("除数不能为0")
	}
	return num1 / num2, nil
}
