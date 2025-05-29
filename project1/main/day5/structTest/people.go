package structTest

import "fmt"

type People struct {
	Name string //首字母大写，外部包可以访问
	Age  int
	Sex  string //首字母小写，外部包不可以访问
}

func (p People) Hello() {
	fmt.Println("hello")
}

func (p People) GetInfo() string {
	return fmt.Sprintf("姓名：%s，年龄：%d，性别：%s", p.Name, p.Age, p.Sex)
}
