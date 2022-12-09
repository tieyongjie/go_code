package main

import "fmt"

//无限循环的控制,直到生成了99这个数
func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break // 默认跳出最近的for循环
			}
			fmt.Println("j=", j)
		}
	}
}
