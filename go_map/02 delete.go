package main

import (
	"fmt"
)

func main() {
	countryCapotalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	for country := range countryCapotalMap {
		fmt.Println(country, "首都是", countryCapotalMap[country])
	}
	delete(countryCapotalMap, "France")
	fmt.Println("删除意大利")
	// 删除之后的地图
	for country := range countryCapotalMap {
		fmt.Println(country, "首都是", countryCapotalMap[country])
	}
}
