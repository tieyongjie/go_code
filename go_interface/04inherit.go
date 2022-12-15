package main

import "fmt"

type Person struct { // 抽象出的字段
	name string
}

func (p *Person) GetName() {
	fmt.Println(p.name)
}

type Inherit struct {
	Person
}

func main() {
	i := Inherit{Person{"火鸡面"}}
	i.GetName()
}
