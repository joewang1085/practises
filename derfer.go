package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 执行时，先将a的值复制一份，然后执行 calc("10", a, b),然后将整个defer内容放到一个“栈”中
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	//执行顺序：   1. 第一个defer里的 calc("10"...)；2.第二个defer里的calc("20"...);3.第二个defer；4.第一个defer
	//原因： defer传参原理：  这里的defer参数都是值类型，所以先将参数的值内容拷贝一份，
	//所以第一个defer执行时：  先拷贝当前的 a,和calc("10",a,b)的的值作为参数，而calc（"10"）是个函数，所以函数栈的形式先执行calc("10")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
