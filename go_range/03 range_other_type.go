package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for index, num := range nums {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}
	//	range 也可以用在 map 的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		// range也可以用来枚举 Unicode 字符串,第一个参数是字符的索引,第二个是字符的unicode值本身
		for i, c := range "go" {
			fmt.Println(i, c)
		}
	}
}
