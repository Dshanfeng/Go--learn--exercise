package initTest

import "fmt"

var InitDemoTest = func() string {
	fmt.Println("初始化initDemoTest.go中全局变量initDemoTest")
	return "initDemoTest"
}()

func init() {
	fmt.Println("正在调用initDemoTest.go的init函数")
}
