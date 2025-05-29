package main

import "fmt"

/*
结构体,是Go语言类似其他语言的面向对象特性的一种实现，结构体可以包含多个字段
结构体的字段可以是任意类型，甚至是结构体本身，通过结构体可以组合出各种复杂的数据结构

结构体是值类型
结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字、个数和类型）
*/
type People struct {
	Name string //首字母大写，外部包可以访问
	Age  int
	sex  string //首字母小写，外部包不可以访问
}

// 工厂方法，也就是构造函数，返回一个结构体的指针，用于创建结构体的实例
// Go语言没有构造函数，但是可以通过工厂方法来实现
// 工厂方法的命名规范：New + 结构体名，首字母大写，返回值是结构体的指针
// 在结构体名首字母为小写时，其他包是不能访问的，但是可以通过工厂方法来创建结构体的实例
func NewPeople(name string, age int, sex string) *People {
	return &People{
		Name: name,
		Age:  age,
		sex:  sex,
	}
}

// 结构体的方法
// 方法是绑定了特定类型的函数，只有通过该类型调用才可以使用，是专门为这种类型服务的函数
// 方法通过将调用者作为一个参数传递给方法来实现类型绑定
// 方法不是结构体独有的，自定义类型也可以定义方法
// 方法的定义格式：func (接收者 类型) 方法名(参数列表) 返回值列表 {方法体}
func (p People) getName() string {
	//结构体的方法可以访问结构体的字段
	return p.Name
}

func (p *People) setName(name string) {
	p.Name = name
}

func (p People) getAge() int {
	return p.Age
}

func (p *People) setAge(age int) {
	p.Age = age
}

func (p People) getSex() string {
	return p.sex
}

func (p *People) setSex(sex string) {
	p.sex = sex
}

func main() {
	//声明一个结构体变量,默认值为零值
	var p People
	fmt.Printf("p的类型：%T， p = %v\n", p, p)

	//给结构体变量的字段赋值
	p.Name = "张三"
	p.Age = 18
	p.sex = "男"
	fmt.Printf("p = %v\n", p)

	//声明结构体变量的同时，给结构体变量的字段赋值
	p1 := People{"李四", 20, "男"}
	//这种方式，必须按照结构体的字段顺序赋值，需给全部字段赋值，否则会报错
	fmt.Printf("p1 = %v\n", p1)

	p2 := People{
		Name: "王五",
		Age:  18,
		sex:  "男",
	}
	//指定字段赋值，不需要按照结构体的字段顺序赋值，可以只给部分字段赋值，其他字段会赋值为零值
	fmt.Printf("p2 = %v\n", p2)

	//使用new函数创建结构体变量，返回的是结构体的指针
	p3 := new(People)
	(*p3).Name = "赵六"
	p3.Age = 18
	p3.sex = "男"
	//Go语言的设计者为了程序员使用方便，底层会对p3.Name进行取值运算，
	//因此可以直接使用p3.Name，而不需要使用(*p3).Name
	fmt.Printf("p3的类型：%T, p3 = %v\n", p3, p3)

	//使用&取结构体变量的地址
	var p4 *People = &People{"田七", 18, "男"}
	fmt.Printf("p4的类型：%T, p4 = %v\n", p4, p4)

	fmt.Println("-----------\n")

	//结构体是值类型，在函数参数传递时，会进行值拷贝，操作副本空间后通过返回值返回
	//所以在函数中修改结构体字段的值，不会影响到原始结构体的值
	//想要直接修改结构体字段的值，需要传递结构体的指针

	//调用者可以是值也可以是指针，底层会根据方法的接收者类型进行转换
	//值接受者方法，会将调用者复制一份，然后传递给方法，方法内部操作的是副本空间
	//指针接受者，会将调用者的地址传递给方法，方法内部操作的是原始空间
	//底层会自动的装箱和拆箱，不需要手动的转换，放心的用，但有一定的性能损耗

	fmt.Println(p4.getName()) //用指针调用值接受者方法
	p.setSex("女")             //用值调用指针接受者方法
	fmt.Println(p.getName())
}
