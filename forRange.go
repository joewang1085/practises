package main

import (
	"fmt"
)

type student struct {
	Age  int
	Name string
}

func main() {
	m := make(map[string]*student)
	sm := make(map[int]string)
	stus := []student{
		{Age: 20, Name: "joe"},
		{Age: 30, Name: "wh"},
	}
	str := []string{
		"helle",
		"hi",
	}

	for _, s := range stus { // 理解for range
		m[s.Name] = &s
	}

	for i, s := range str { // 理解for range
		sm[i] = s
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, v := range sm {
		fmt.Println(k, v)
	}

}
