package main

import "fmt"

type Order struct {
	number  int
	address string
	price   float32
}

func main() {
	var order Order
	order.number = 10
	order.price = 15.6
	order.address = "shenzhen"
	fmt.Println(order)
}
