package main

import "fmt"

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	orderPrice := totalPrice(5)
	fmt.Println(orderPrice(5))
	fmt.Println(orderPrice(5))
	fmt.Println(orderPrice(5))
}
