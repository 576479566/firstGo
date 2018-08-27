package main

import (
	"fmt"
	"time"
)

var c, python, java bool

func main() {
	//:= 声明并赋值,:=不能写在函数外，函数外的必须以关键字开始，var func等等
	//常量不能使用 := 语法定义。
	var i int
	//自动获得类型
	var str = "no!"
	fmt.Println("1")
	a,b := swap("hello","world")
	fmt.Println(a,b)
	fmt.Println(split(17))
	fmt.Println(c,python,i)
	fmt.Println(str)
	time.Sleep(5 * time.Second)

}

func swap(x,y string)(string,string)  {
	return y,x
}

//无参的return返回的就是x和y（的当前值，按名字
func split(sum int)(x,y int)  {
	x = sum*4/9
	y = sum-x
	return
}