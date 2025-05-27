package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	//切片，是Go语言特有的数据结构，建立在数组类型之上，可以看作可变长度的数组
	//切片的操作
	var slice = make([]int, 5)
	fmt.Println(slice)
	fmt.Printf("%T\n", slice)

	//切片的增改，就是类似数组的操作
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	fmt.Println(slice)

	slice[2] = 0
	fmt.Println(slice)

	//拼接,就是底层创建一个新数组，将你拼接的内容复制到新数组中，然后将新数组的地址赋值给切片
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)

	//切片删除元素才是难点
	//通过拼接新切片来实现，将要删除的元素剔除后，再拼接新切片
	var err error
	slice, err = deleteElement(slice, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)

	//或将切片指向你创建的新数组(append)，其实两者本质上是一致，都是创建了新的数组
	slice, err = deleteElement2(slice, 1, 5, 3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice)
}

func deleteElement(slice []int, index int) ([]int, error) {
	if index < 0 || index >= len(slice) {
		return slice, errors.New("数组越界异常")
	}

	return append(slice[:index], slice[index+1:]...), nil
}

//func deleteElement2(slice []int, index ...int) ([]int, error) {
//	for i := 0; i < len(index); i++ {
//		if index[i] < 0 || index[i] >= len(slice) {
//			return slice, errors.New("数组越界异常")
//		}
//	}
//
//	var newSlice = make([]int, len(slice)-len(index))
//	n, flag := 0, true
//	for i := 0; i < len(slice); i++ {
//		flag = true
//
//		for j := 0; j < len(index); j++ {
//			if i == index[j] {
//				flag = false
//				index = append(index[:j], index[j+1:]...)
//				break
//			}
//		}
//
//		if flag {
//			newSlice[n] = slice[i]
//			n++
//		}
//	}
//	return newSlice, nil
//}

// AI改动
// deleteElement2 从切片中删除多个指定索引位置的元素。
// 参数 slice 是要操作的切片，index 是可变参数，包含多个要删除元素的索引。
// 返回删除元素后的新切片和可能出现的错误。
func deleteElement2(slice []int, index ...int) ([]int, error) {
	// 检查所有要删除的索引是否越界
	for i := 0; i < len(index); i++ {
		if index[i] < 0 || index[i] >= len(slice) {
			return slice, errors.New("数组越界异常")
		}
	}

	// 为避免重复删除，先对索引进行排序并去重
	// 对索引进行排序
	sort.Ints(index)
	// 去重,这个去重挺有意思的，debug研究一下，先排序重复数字就会相邻，遍历比较相邻数字是否相等，通过i和j两个指针来比较
	// i是快指针，j是慢指针，两者相差1，当j和i所指向数不相等时，j++i++,两者还是i比j大1
	// 但当j和i所指向数相等时,重复数字找到了,i++，j不动，此时两者i比j大2，再次比较i和j所指向的数是否相等
	// 实现了跳过重复数字的效果，此时不是重复数字，j++，就会将i指向数字覆盖到重复数字的位置，i++，i和j又相差2

	//遇到连续重复数字，就加大i和j的差值，跳过重复数字，直到i和j所指向的数字不相等，再将i所指向的数字覆盖到j+1的位置
	//实现不相邻数字不重复，最后将j后面没用的数字删除，就实现了去重
	j := 0
	for i := 1; i < len(index); i++ {
		if index[i] != index[j] {
			j++
			index[j] = index[i]
		}
	}
	index = index[:j+1]

	var newSlice = make([]int, len(slice)-len(index))
	n := 0
	delIndex := 0
	for i := 0; i < len(slice); i++ {
		// 如果当前索引等于要删除的索引，则跳过该元素
		//delIndex < len(index),自己就是一开始没想到这条语句写在哪里放弃了这种方案，虚心学习
		if delIndex < len(index) && i == index[delIndex] {
			delIndex++
			continue
		}
		newSlice[n] = slice[i]
		n++
	}
	return newSlice, nil
}
