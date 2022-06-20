package main

import "fmt"

//题一
//func main() {
//	var a int = 300
//	//*int 用来接收指针变量
//	var ptr *int = &a
//	fmt.Println("ptr:",ptr)
//}
//题二
//func main() {
//	var a int = 300
//	//*int 用来接收指针变量
//	var ptr *float = &a
//	//类型错误
//	fmt.Println("ptr:",ptr)
//}
// 题三
func main() {
	var a int = 100
	var b int = 400
	var ptr *int = &a
	*ptr = 100
	ptr = &b
	*ptr = 200
	fmt.Printf("a=%d,b=%d,*ptr=%d", a, b, *ptr)

}
