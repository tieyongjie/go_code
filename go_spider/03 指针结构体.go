package main

import "fmt"

// Person 结构体
type Person struct {
	name string
	city string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// Dream Person 做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言", p.name)
}

func main() {
	p1 := NewPerson("white.tie", 25)
	p1.Dream()
}
