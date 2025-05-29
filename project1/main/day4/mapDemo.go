package main

import "fmt"

func main() {
	//map引用数据类型，地址传递，直接修改map本身,
	//无序去重的特点，不能通过普通for遍历
	//map存储的是键值对，键key,值value
	//key不能重复，类型必须是能比较的类型，如：int、float、bool、string

	//声明
	var aMap map[string]int //未分配空间
	//[]括号内是key的类型，[]括号外是value的类型
	//key的类型必须是能比较的类型，如：int、float、bool、string
	//value的类型可以是任意类型

	//创建对象
	//内存：堆
	aMap = make(map[string]int, 10) //可指定长度，也可以不指定
	//make创建的map对象，其内存是在堆中，底层是在堆中先创建了一个数组，数组长度便是map的长度，然后创建map指针指向数组
	fmt.Printf("aMap的类型：%T, aMap=%v\n", aMap, aMap)

	//内存：栈
	bMap := map[string]int{
		"张三": 18,
		"李四": 20,
	}
	fmt.Printf("bMap的类型：%T, bMap=%v, bMap的长度：%d\n", bMap, bMap, len(bMap))

	fmt.Println("---------\n")

	//增删改查

	//增和改，类似数组的操作方式，索引是key,值是value
	//map有去重的功能，key不存在是添加，key存在会覆盖键值对的值
	aMap["王五"] = 20
	aMap["周六"] = 18
	fmt.Println(aMap)

	fmt.Println("--------\n")

	aMap["王五"] = 19
	fmt.Println(aMap)

	fmt.Println("--------\n")

	//删
	delete(aMap, "王五")
	//若是aMap中不存在key也不会报错，什么也不处理
	fmt.Println(aMap)

	fmt.Println("--------\n")

	//查
	value, contain := aMap["王五"]
	if contain {
		fmt.Println("aMap中存在王五:", value)
	} else {
		fmt.Println("aMap不存在王五")
	}

	fmt.Println("---------\n")

	//遍历 for-range
	for k, v := range aMap {
		fmt.Println(k, "=", v)
	}

	//清空内容
	//  遍历一一删除
	//  让对象引用新的空间，旧空间没有对象引用，空间就会被回收自动释放空间
	for k, _ := range aMap {
		delete(aMap, k)
	}

	bMap = make(map[string]int, 0)

	fmt.Printf("bMap的长度：%d\n", len(bMap))
	fmt.Printf("aMap的长度：%d\n", len(aMap))
}
