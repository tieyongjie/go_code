package main

import "fmt"

// 打印半个个金字塔
//func main() {
//	for i := 0; i < 5; i++ {
//		for j := 0; j <= i; j++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}

// 打印完整的金字塔
func main() {
	var totallevel int = 20
	for i := 1; i < totallevel; i++ {
		for j := 1; j <= totallevel-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totallevel-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
