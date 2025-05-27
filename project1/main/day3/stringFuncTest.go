package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//不用导包就能使用的函数是内置函数，定义在builtin包下

	//1、统计字符串长度，以字节为单位
	str1 := "Hello!你好！"
	fmt.Println("str1的长度：", len(str1))

	//2、遍历字符串
	for i, v := range str1 {
		fmt.Printf("索引：%d,值：%c\n", i, v)
	}

	fmt.Println("----------")

	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Println("索引：", i, " 值：", string(r[i]))
	}
	//从索引上看就可以看出，字符串是以字节为单位的，每个汉字都占据了3个索引
	fmt.Println("---------")

	//3、字符串转整数
	num, _ := strconv.Atoi("888")
	fmt.Println(num)

	//4、整数转字符串
	str2 := strconv.Itoa(num)
	fmt.Println(str2)

	str2 = fmt.Sprintf("%d", num)
	fmt.Println(str2)
	fmt.Println("---------")

	//5、统计字符串中是否包含字串
	str3 := "golanggojavagolanggo"
	fmt.Println(strings.Contains(str3, "go"))

	//6、包含多少个子串
	fmt.Println(strings.Count(str3, "go"))
	fmt.Println("----------")

	//7、字符串比较，
	//字符串属于基础数据类型，直接比较即可
	fmt.Println(str3 == "golanggojava")
	//不区分大小写
	fmt.Println(strings.EqualFold(str3, "GolangGoJavaGolangGO"))
	fmt.Println("----------")

	//8、返回字符串中第一个子串的位置，如果没有返回-1
	fmt.Println("str3的第一个go字串的索引：", strings.Index(str3, "go"))
	fmt.Println("--------")

	//9、字符串替换
	//n指定替换的个数，-1是将匹配的子串全部替换
	fmt.Println(strings.Replace(str3, "go", "C", -1))
	fmt.Println("--------")

	//10、按照指定字符分割标识，返回字符串数组
	strs := strings.Split(str3, "go")
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
	fmt.Println("--------")

	//11、大小写转换
	fmt.Println(strings.ToLower(str3))
	fmt.Println(strings.ToUpper(str3))
	fmt.Println("--------")

	//12、将字符串左右两边的空格去除
	fmt.Println(strings.TrimSpace("    abc    "))

	//13、将字符串左右两边的指定字符去除
	fmt.Println(strings.Trim(str3, "golanggo"))

	//14、将字符串左边的指定字符去除
	fmt.Println(strings.TrimLeft(str3, "golang"))

	//15、将字符串右边的指定字符去除
	fmt.Println(strings.TrimRight(str3, "go"))
	//有贪狼多匹配的特性，匹配上的内容有相邻的相同字母也会被去除
	fmt.Println("-----------")

	//16、判断字符串是不是以指定字符开头
	if strings.HasPrefix(str3, "golang") {
		str3 = strings.TrimLeft(str3, "golang")
	}
	fmt.Println(str3)

	//17、判断字符串是不是以指定字符结尾
	if strings.HasSuffix("abc.txt", ".txt") {
		fmt.Println("文本文件")
	}

}
