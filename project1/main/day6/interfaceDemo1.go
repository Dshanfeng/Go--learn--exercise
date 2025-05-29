package main

import "fmt"

/*
接口是一种抽象类型，它定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
接口是方法签名的集合，不包含方法的具体实现，仅定义应该做什么
*/

type Live interface {
	eat()
	sleep()
}

/*
接口的实现规则：
1、隐式实现：

	一个对象只要全部实现了接口中的方法，那么就实现了这个接口，不用显示的声明实现了哪个接口

2、接口值：

	接口变量的底层是一个结构体，结构体中有两个字段，一个是类型，一个是值
	接口值是由一个具体类型和该类型的值两部分组成的，接口值可以用来调用接口中的方法
	接口变量可以存储任何实现了该接口的类型的值，体现了多态的特性
*/
type Dog struct {
	name    string
	variety string
}

// 实现的方法必须完全符合接口的方法标签，才能实现接口
func (d Dog) eat() {
	fmt.Println("狗", d.name, "正在吃骨头")
}
func (d Dog) sleep() {
	fmt.Println("狗", d.name, "正在睡觉")
}
func (d Dog) getInfo() string {
	return fmt.Sprintln("狗的名字是：", d.name, "品种是：", d.variety)
}

type Cat struct {
	name    string
	variety string
}

func (c Cat) eat() {
	fmt.Println("猫", c.name, "正在吃鱼")
}
func (c Cat) sleep() {
	fmt.Println("猫", c.name, "正字睡觉")
}
func (c Cat) getInfo() string {
	return fmt.Sprintln("猫的名字是：", c.name, "品种是：", c.variety)
}

/*
接口的特性：
1、抽象性：
     接口只关心对象的行为，不关心对象的具体实现，它只注重方法不关心具体类型，任何类型都可以实现接口，只要该类型实现了接口中的所有方法
2、空接口：
     空接口是没有任何方法的接口，它可以表示任何类型，任何类型都可以实现空接口，类似java的Object类可以接受任何类型，包括基本类型和自定义类型
3、接口嵌套：
     接口可以嵌套，一个接口可以包含另一个接口，这样就可以实现接口的继承，接口的继承可以实现代码的复用，避免重复的代码
     这点需要注意，Go语言更换注重简洁性，小接口组合比大接口嵌套更简洁，所以Go语言不建议接口嵌套
组合优于继承，Go语言提倡使用组合来实现代码的复用，而不是继承
实现 松耦合，高内聚 的设计
*/

type Animal interface {
	Live
	//注意：接口嵌套不能有同名的方法，否则会报错，有同名的方法，编译器就无法分辨，重复定义肯定会报错
	//想要实现接口的方法，需要实现接口的所有方法，否则会报错
	run()
	testValue(name string)
}

type People struct {
	name string
	sex  string
	age  int
}

func (p People) eat() {
	fmt.Println("人", p.name, "正在吃")
}
func (p People) sleep() {
	fmt.Println("人", p.name, "正在睡觉")
}
func (p People) run() {
	fmt.Println("人", p.name, "正在两条腿直立奔跑")
}
func (p People) testValue(name string){
	p.name = name
}

func main() {
	var dog Dog = Dog{"大黄", "土狗"}
	//Go语言的多态性是通过接口来实现的，接口变量可以接受任何实现了该接口的类型的值，通过接口变量可以调用实现类型绑定的接口方法
	var l Live = dog
	l.eat()
	l.sleep()

	fmt.Println("------------")

	cat := Cat{"阿花", "狸花猫"}
	test(cat)

	fmt.Println("------------")

	isClass(cat)

	fmt.Println("------------")

	p := People{"张三", "男", 18}
	var a Animal = p //p实现了Animal接口的所有方法，所以p可以赋值给a
	var l1 Live = p  //Animal接口嵌套了Live接口，所以p也可以赋值给l
	var l2 Live = a  //也可以将a赋值给l,你实验a的类型发现a本质上还是Animal类型，Animal类型实现了Live接口的所有方法，所以a可以赋值给l
	fmt.Println(l1, l2)
	fmt.Println()

	//引出指针值底层的原理：
	//1、接口类型是引用类型，接口变量存储的是一个指向具体值的指针
	var l3 Live //未赋值的接口变量值nil
	fmt.Println(l3)
	fmt.Println()

	//2、接口值是由一个具体类型和该类型的值两部分组成的，接口值可以用来调用接口中的方法
	//将一个值类型赋值给一个接口变量时，Go语言会拷贝该值的副本到接口变量中并存储值的具体类型信息，返回指向该副本的指针
	//所以对接口变量的修改不会影响到原始值
	a.testValue("李四")
	fmt.Println("a:", a)
	fmt.Println("p:", p)
	fmt.Println("l1:", l1)
	fmt.Println()

}

func test(l Live) {
	l.eat()
	l.sleep()
	//通过接口变量只能调用接口方法，不能调用类型的方法
	//fmt.Println(l.getInfo())
}

func isClass(l Live) {
	fmt.Printf("%T, %v\n", l, l)
	//类型断言    类型断言是一种将接口类型转换为具体类型的操作
	//类型断言的语法是：value, ok := interfaceVar.(Type)
	//value是转换后的具体类型的值，ok是一个布尔值，表示转换是否成功
	//如果转换失败，ok的值为false，value的值为Type类型的零值
	//如果转换成功，ok的值为true，value的值为转换后的值
	switch l.(type) {
	case Dog:
		fmt.Println("这是一个Dog类型")
	case Cat:
		fmt.Println("这是一个Cat类型")
	}

	fmt.Println()

	if d, ok := l.(Dog); ok {
		fmt.Println("这是一个Dog类型")
		fmt.Printf("%T, %v\n", d, d)
		fmt.Println(d.getInfo())
	}

	if c, ok := l.(Cat); ok {
		fmt.Println("这是一个Cat类型")
		fmt.Printf("%T, %v\n", c, c)
		fmt.Println(c.getInfo())
	}
}
