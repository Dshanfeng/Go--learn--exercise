package main

import (
	"fmt"
	"project1/main/day5/structTest"
)

// 结构体的嵌入,是Go语言实现代码复用的一种方式，结构体的嵌入可以将一个结构体嵌入到另一个结构体中，
// 被嵌入的结构体称为嵌入字段，嵌入字段可以是结构体本身，也可以是结构体的指针
// 结构体的嵌入是指将一个结构体嵌入到另一个结构体中，被嵌入的结构体称为嵌入字段，嵌入字段可以是结构体本身，也可以是结构体的指针
// 嵌入字段的字段可以直接被外部结构体访问，不需要通过嵌入字段的名称来访问

type Student struct {
	structTest.People //嵌入结构体直接写结构体的名称，匿名嵌入
	//只能访问嵌入结构体的导出字段，不能访问嵌入结构体的非导出字段
	School string
	Class  string
}

type Teacher struct {
	*structTest.People //嵌入结构体的指针，匿名嵌入
	School             string
	Course             string
}

type Parents struct {
	fatcher structTest.People
	mather  structTest.People
}

type Graduate struct {
	Student //多重嵌入
	Parents
}

func (stu Student) getInfo() string {
	return fmt.Sprintf("学校：%s，班级：%s，%s\n", stu.School, stu.Class, stu.People.GetInfo())
}

func (p Parents) getInfo() string {
	return fmt.Sprintf("父亲：%s，年龄：%d，性别：%s\n母亲：%s，年龄：%d，性别：%s\n",
		p.fatcher.Name, p.fatcher.Age, p.fatcher.Sex, p.mather.Name, p.mather.Age, p.mather.Sex)
}

// 实现代码复用的方式就是，用结构体中的嵌入字段来调用嵌入结构体的方法来实现代码复用
// 类似于方法的重写
func (g Graduate) getInfo() string {
	return fmt.Sprintln(g.Student.getInfo(), g.Parents.getInfo())
}

func main() {
	//创建对象时，嵌入字段的字段可以直接赋值
	var stu1 Student = Student{structTest.People{Name: "张三", Age: 18, Sex: "男"}, "清华大学", "三年一班"}
	teacher1 := Teacher{&structTest.People{Name: "李四", Age: 30, Sex: "女"}, "北京大学", "数学"}

	//访问嵌入字段的字段
	//嵌入字段的字段可以直接访问，不需要通过嵌入字段的名称来访问
	fmt.Println(stu1.Name, stu1.Age, stu1.Sex)
	fmt.Println(teacher1.Name, teacher1.Age, teacher1.Sex)

	//调用嵌入字段的方法
	stu1.Hello()
	teacher1.Hello()

	fmt.Println("-------------")

	//都可以直接访问嵌入结构体的东西，不需要通过嵌入字段的名称来访问，因为底层会做判断
	//访问字段会先在当前结构体中查找，如果没有找到，会在嵌入字段中查找，如果还没有找到，会在嵌入字段的嵌入字段中查找，以此类推
	//访问方法也会先在当前结构体中查找，如果没有找到，会在嵌入字段中查找，如果还没有找到，会在嵌入字段的嵌入字段中查找，以此类推
	//这是在嵌入字段与结构体字段没有重复同名的情况下
	//如果嵌入字段与结构体字段有重复同名或多重嵌入下的字段重复，都会出现冲突
	//需要使用嵌入字段的名称来访问，也可以不适用匿名嵌入，为其命名，这样就可以避免冲突

	graduate := Graduate{
		Student: Student{structTest.People{Name: "王五", Age: 20, Sex: "男"}, "清华大学", "三年一班"},
		Parents: Parents{
			fatcher: structTest.People{Name: "赵六", Age: 40, Sex: "男"},
			mather:  structTest.People{Name: "钱七", Age: 40, Sex: "女"},
		},
	}
	//访问字段
	//访问字段是类似就近原则的，先在当前结构体中查找，如果没有找到，会在嵌入字段中查找，如果还没有找到，会在嵌入字段的嵌入字段中查找，以此类推
	//在本例中,Graduate结构体中没有Name字段，所以会在Student结构体中查找，所以直接访问graduate.Name访问的就是Student结构体中的Name字段
	fmt.Printf("大学生：%s，年龄：%d，性别：%s\n", graduate.Name, graduate.Age, graduate.Sex)
	//fmt.Printf("大学生：%s，年龄：%d，性别：%s\n", graduate.Student.Name, graduate.Student.Age, graduate.Student.Sex)
	fmt.Printf("父亲：%s，年龄：%d，性别：%s\n", graduate.fatcher.Name, graduate.fatcher.Age, graduate.fatcher.Sex)
	fmt.Printf("母亲：%s，年龄：%d，性别：%s\n", graduate.mather.Name, graduate.mather.Age, graduate.mather.Sex)

	//访问方法
	//方法都是绑定了类型的，通过什么类型调用，就会调用相应的方法，所以在本例中，graduate.GetInfo()调用的是Graduate结构体中的GetInfo方法
	graduate.GetInfo()

	/*
			总结：
				嵌入结构体实现了代码复用，也变相的实现了继承，这也是Go语言继承中关键的一部分
		        其原理很简单，就是一个个的变量，这些变量有基本数据类型也有结构体复杂数据类型，画画内存图就可以理解了，并不复杂
		    注意事项：
		         存在重名命的字段或方法时，需要使用嵌入字段的名称来访问，想要细致的区分可以为嵌入字段命名，这样就可以避免冲突
	*/
}
