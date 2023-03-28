package testfmt

import (
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Print("古藤老树昏鸦\n")
	fmt.Printf("我是%s\n", "巴啦啦小魔仙")
	fmt.Println("\tdyy")
}
func TestFprint(t *testing.T) {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		fileObj.Close()
		return
	}
	name := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
	fileObj.Close()
}
func TestSprint(t *testing.T) {
	s1 := fmt.Sprint("枯 藤", 123, "xx")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}
func TestErrorf(t *testing.T) {
	err := fmt.Errorf("这是一个错误")
	fmt.Println(err)
}
func TestScan(t *testing.T) {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
