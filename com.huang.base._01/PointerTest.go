package main

import "fmt"

func main() {
	a := 1
	b := 2
	changeNum2(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func changeNum(num1 int, num2 int) {
	num1, num2 = num2, num1
}

func changeNum2(num1 *int, num2 *int) {
	*num1, *num2 = *num2, *num1
}
