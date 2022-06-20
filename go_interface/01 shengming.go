package main

import "fmt"

type usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}
type Camera struct {
}

func (phone Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (phone Phone) Stop() {
	fmt.Println("手机结束工作...")
}

// 让相机 实现 usb接口的方法

func (phone Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (phone Camera) Stop() {
	fmt.Println("相机结束工作...")
}

// 编写一个方法 Working 方法，接收一个usb接口类型变量
// 只要是实现了 usb接口（所谓实现接口）
type Computer struct {
}

func (c Computer) Working(usb usb) {
	// 通过usb接口来调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
