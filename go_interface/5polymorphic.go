package main

import "fmt"

type Birds interface {
	fly()
}

type Dove struct {
}

type Eagle struct {
}

func (d *Dove) fly() {
	fmt.Println("鸽子在飞")
}

func (e *Eagle) fly() {
	fmt.Println("鹰在飞")
}

func main() {
	var b Birds
	dove := Dove{}
	eagle := Eagle{}
	b = &dove
	b.fly()
	b = &eagle
	b.fly() // 鹰在飞

}
