package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages       map[string]int
	sync.Mutex // 多线程的map ,为了安全可以加锁
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	//	ua.Lock()
	//	defer ua.Unlock() //重点：这里使用的同步锁,不止“写”需要加锁，读也需要“加锁”，否则正在“写”的时候进行“读”就会报错
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	fmt.Println("map线程安全")
}
