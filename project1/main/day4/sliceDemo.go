package main

import "fmt"

func main() {
	//切片，是Go语言特有的数据结构，建立在数组类型之上，可以看作可变长度的数组
	//切片是数组的一个视图，属于引用数据类型，是一个结构体，本身不存储数据结构而是指向数组

	//结构体由三部分组成：指针、长度、容量
	//指针指向被引用数组的第一个元素，存储其地址
	//长度是切片的长度，一般为引用片段的长度，不能超过容量
	//容量是切片的容量，容量不需要太过在意，切片底层有扩容机制，当长度超过容量时，会自动扩容
	//切片的长度和容量都是可以动态变化的，切片的长度不能大于容量，切片的容量不能小于长度

	//切片的定义
	//1、引用已存在数组的片段，片段可以是整个数组，片段范围是包左不包右的
	var intArr = [5]int{1, 2, 3, 4, 5}
	var slice1 = intArr[1:3] //var 切片名 = 数组名[起始索引:终止索引]
	fmt.Printf("slice1的类型：%T, slice1 = %v, &slice1 = %p\n", slice1, slice1, &slice1)
	fmt.Printf("&intArr[1] = %p\n", &intArr[1])
	fmt.Printf("slice1的长度：%d, slice1的容量：%d\n", len(slice1), cap(slice1))
	//slice1的结构体中指向底层数组的指针是密封的我们无法通过直接访问其字段
	//我们引用栈中的数组，其创建的切片也存储于栈中

	fmt.Println("----------\n")

	//引用数组的简写方式
	//起始是0，结束是len可以省略不写
	slice1 = intArr[:3]
	fmt.Printf("slice1的长度：%d, slice1的容量：%d\n", len(slice1), cap(slice1))
	fmt.Println("范围：0~2", slice1)

	slice1 = intArr[1:]
	fmt.Printf("slice1的长度：%d, slice1的容量：%d\n", len(slice1), cap(slice1))
	fmt.Println("范围：1~5", slice1)

	slice1 = intArr[:]
	fmt.Printf("slice1的长度：%d, slice1的容量：%d\n", len(slice1), cap(slice1))
	fmt.Println("范围：0~5", slice1)

	//这种引用切片容量的计算公式：
	//          容量 = 源长度 - 起始索引
	//          源：可以是数组也可以是切片，源长度指的就是引用物的长度
	sliceSon := slice1[1:3]
	fmt.Printf("sliceSon的长度：%d, sliceSon的容量：%d, sliceSon = %v\n",
		len(sliceSon), cap(sliceSon), sliceSon)

	fmt.Println("---------\n")

	//介绍一下切片的两种特性，扩容 和 共享数组

	//扩容
	//切片长度就是结束索引到起始索引，也就是存有值的哪部分
	//而切片可以看作可变长度的数组，也就是说切片可以动态增长append内置函数，容量在这里就有作用了
	fmt.Println("添加元素前：")
	fmt.Printf("sliceSon的长度：%d, sliceSon的容量：%d，sliceSon = %v\n",
		len(sliceSon), cap(sliceSon), sliceSon)

	sliceSon = append(sliceSon, 666, 888) //拼接
	fmt.Println("添加元素后：")
	fmt.Printf("sliceSon的长度：%d, sliceSon的容量：%d，sliceSon = %v\n",
		len(sliceSon), cap(sliceSon), sliceSon)

	//共享数据也很好理解，因为切片本身并不是储存数据的，只是引用了数组，
	//所以对切片数据进行修改，同时它引用的源也会被修改数据
	//在这里，sliceSon引用了slice1范围1-2，我在sliceSon中添加元素，slice1中2索引后的元素就会被覆盖
	//而slice1引用了intArr数组，intArr数组也受到了影响，所以切片修改数据要慎重
	fmt.Println("slice1切片：", slice1)
	fmt.Println("intArr数组；", intArr)
	//数据共享的特性我们可以将切片看为数组的视图

	fmt.Println("---------\n")

	//当前长度与容量已经相同，此时在添加元素就会扩容
	sliceSon = append(sliceSon, 100) //拼接
	fmt.Println("添加元素后：")
	fmt.Printf("sliceSon的长度：%d, sliceSon的容量：%d，sliceSon = %v\n",
		len(sliceSon), cap(sliceSon), sliceSon)
	//观察引用对象的影响
	fmt.Println("slice1切片：", slice1)
	fmt.Println("intArr数组；", intArr)
	//我们又会发现，容量扩大了，引用源的数据这回没受到影响
	//到扩容时，底层会创建一个新的数组，将切片引用的老数组片段复制到新数组中，在添加元素
	//新数组是隐藏的我们是无法直接操作的，只能通过切片间接的操作
	//这种拥有独立隐藏的私有数组的切片可以直接看成一个可变数组，一个存储数据的数据结构

	fmt.Println("----------------\n")

	//2、make动态分配内存来创建切片
	var slice2 = make([]int, 5, 7)
	//三个参数，切片类型，切片长度，切片容量
	//   不写容量时，默认长度便是其容量
	//make创建的切片其内存是在堆中，底层是在队中先创建了一个数组，数组长度便是切片长度，然后创建切片指针指向数组
	//这个底层创建的数组是隐藏的，我们也是不能直接操作的，需通过slice2切片来间接的访问操作
	//这种的切片我们更可以看成是一个可变长度的数组
	fmt.Printf("slice2的类型：%T, slice2 = %v, &slice2 = %p\n", slice2, slice2, &slice2)
	fmt.Printf("slice2的长度：%d, slice2的容量：%d\n", len(slice2), cap(slice2))

	fmt.Println("----------------\n")

	//3、直接指定无名数组，这种更类似make创建切片
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)

	fmt.Println("---------\n")

	//空切片和nil切片是有去别的，就像空串和null字符串一样
	//空切片有一个长度为零的数组
	//nil底层没有数组，长度为nil
	var slice4 []int = nil
	fmt.Println(slice4)

	//总结：
	//   1、引用数组或切片创建的切片，将其看成是源的视图，有数据共享的特点，要谨慎的修改数据
	//   2、make动态分配的、直接指定无名数组的、经历过扩容的都属于拥有独立隐藏的私有数组的切片
	//     可以直接看成一个可变数组，一个存储数据的数据结构
	//   3、切片结构体的三部分，指向底层数组的指针
	//                     切片的长度，访问切片时不要超过长度
	//                     切片的容量，有自动扩容存在容量可以不用太在意，只需注意修改数据时是否已经扩容过，是否会影响引用空间
	//   4、应用场景：
	//         动态数据集合、函数参数传递（地址传递可以直接修改数据）
	//         字符串处理（字符串底层是个[]byte，先将字符串转成[]byte便可通过切片引用直接修改字符串）

}
