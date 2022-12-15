package main

import "fmt"

type Person struct {
	name string
}

func (person *Person) setName(name string) {
	person.name = name
}

func (person *Person) Getinfo() {
	fmt.Println(person.name)
}

func main() {
	p := Person{"xcc的螺丝钉"}
	p.Getinfo()
	p.setName("这是个新名字")
	p.Getinfo()
}
