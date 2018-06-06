package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	fmt.Println("interface  == nil?坑")

	if live() == nil { // 带有方法集的接口：data为nil,但itab指明了类型为*Student，且itab包含了*Student的方法集的实现，不为nil
		fmt.Println("AAAAAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	var peo People
	if peo == nil { //带有方法集的接口：data为nil，itab未指明了结构体类型,所以interface 为nil
		fmt.Println("AAAAAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

}

////重点：解析
//接口类型 ： 空接口 和 有方法集合接口
//空接口  var in interface
//有方法集接口  type People interface { Show() }
//底层结构：
//type eface struct{  //空接口
//	_type *_type  // 类型信息
//	data unsafe.Pointer //只向数据的指针；      其中unsafe.Pointer 类似c中void
//}

//type iface struct{  // 带有方法集的接口
//	tab *itab   // 存储类型信息还有结构实现方法的集合
//	data unsafe.Pointer //只向数据的指针
//}

//type _type struct{
//	size uintptr   // 类型大小
//	ptrdata uintptr //前缀持有所有指针的内存大小
//	hash uint32   // 数据hash值
//	tflag tflag  //
//	align uint8  //对齐
//	fieldalign uint8 //嵌入结构体时的对其
//	kind uint8 //
//	alg *typeAlg //函数指针数组，类型实现所有的方法
//	gcdata *byte //
//	str nameOff//
//	ptrToThis typeOf//
//}

//type itab struct{
//	inter *interfacetype//接口类型
//	_type *_type//结构类型
//	link *itab//
//	bad int32//
//	inhash int32//
//	func [1]uintptr//   方法集合
//}

//总结：
//1.iface比eface中间多了一层itab结构。itab存储了——type信息和[]fun方法集。
//2.data只向nil并不代表interface是nil,只有data和 _type(/itab)都是nil,该interface才是nil;
//3.空接口： 只有data 和type同时为空时，interface是nil
//4.带有方法集合的接口 ： 只有data 和type同时为空时,itab未包含了结构体fun方法集的定义时,为nil
