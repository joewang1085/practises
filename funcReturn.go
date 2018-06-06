package main

import (
	"fmt"
)

func main() {

	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))

}

func DeferFunc1(i int) (t int) { //  t的命名使得t的作用域为整个函数 包括defer内部
	t = i
	defer func() { //defer在执行 return t之后执行：  先return t(t=1) ,再执行defer ,因为t的作用域为整个函数所以defer修改了t的值，使得t最终为4
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i         //  t的命名使得t的作用域为函数内部，不包括defer
	defer func() { //defer在执行 return t之后执行：  先return t(t=1) ,再执行defer ,因为t的作用域只为函数内部所以defer无法修改t的值，使得t最终为1
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) { //  t的命名使得t的作用域为整个函数 包括defer内部
	defer func() { //defer在执行 return t之后执行：  先return 2,使得t=2 ,再执行defer ,因为t的作用域为整个函数所以defer修改了t的值，使得t最终为3
		t += i
	}()
	return 2
}
