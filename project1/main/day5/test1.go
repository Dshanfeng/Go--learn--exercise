package main

import (
	"fmt"
	"project1/main/day5/structTest"
)

func main() {
	circle1 := structTest.NewCircle(5)
	//首字母小写的字段，外包是无法直接获取与操作的，只能通过方法来获取与操作
	fmt.Println("circle1的半径为：", circle1.GetRadius())

	fmt.Println("---------------\n")

	circle2 := &structTest.Circle{}
	//无法将值赋给未导出的字段
	circle2.SetRadius(8)

	circle3 := structTest.NewCircle(6.5)
	circle4 := structTest.NewCircle(4.5)

	cSlice := []*structTest.Circle{circle1, circle2, circle3, circle4}
	for k, v := range cSlice {
		fmt.Printf("circle%d半径为：%.2f，面积为：%.2f\n", k+1, v.GetRadius(), v.GetArea())
	}

	fmt.Println("---------------\n")

	//如果写了String方法，那么fmt的输出方法就会自动调用
	for k, v := range cSlice {
		fmt.Printf("circle%d%v", k+1, v)
	}
}
