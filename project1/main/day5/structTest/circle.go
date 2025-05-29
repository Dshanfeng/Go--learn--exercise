package structTest

import "fmt"

var PI = 3.14

/*
封装：通过命名规范和访问控制，将数据和方法封装在一个结构体中，实现数据的隐藏和保护
    1.命名规范：
        1.结构体名首字母大写，外部包可以访问
        2.结构体名首字母小写，外部包不可以访问
    2.访问控制：
        1.结构体的字段首字母大写，外部包可以访问
        2.结构体的字段首字母小写，外部包不可以访问
    3.工厂方法：
        1.工厂方法的命名规范：New + 结构体名，首字母大写，返回值是结构体的指针
        2.在结构体名首字母为小写时，其他包是不能访问的，但是可以通过工厂方法来创建结构体的实例
    4.方法：
        1.方法是绑定了特定类型的函数，只有通过该类型调用才可以使用，是专门为这种类型服务的函数
        2.方法通过将调用者作为一个参数传递给方法来实现类型绑定
        3.方法不是结构体独有的，自定义类型也可以定义方法
        4.方法的定义格式：func (接收者 类型) 方法名(参数列表) 返回值列表 {方法体}
        5.结构体的方法可以访问结构体的字段
*/

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	//可以通过方法来实现数据的验证，保证数据的有效性安全性
	if radius < 0 {
		fmt.Println("半径不能为负数")
		return
	}
	c.radius = radius
}

/*
方法与函数的区别：
	1.方法是作用于特定类型的函数，函数是作用于所有类型的
      两者的调用方式不同，方法是通过类型的实例来调用的，函数是通过函数名来调用的
          函数： 函数名(参数列表)
          方法： 变量.方法名(参数列表)
    2、方法的调用者可以是值类型也可以是指针类型，会根据方法的接收者类型自动对调用者的类型进行转换
      函数的参数必须严格按照定义的参数列表来传递
*/

func (c Circle) GetArea() float64 {
	return PI * c.radius * c.radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("半径为：%.2f，面积为：%.2f\n", c.radius, c.GetArea())
}
