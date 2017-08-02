package practice

import "fmt"

type interfaceOne interface {
	methodOne()
	methodTwo()
}

type interfaceTwo interface {
	methodTwo()
	methodThree()
}

//=================================================================

type one int

func (one one) methodOne() {
	fmt.Println("type one invoke methodOne...")
}

func (one one) methodTwo() {
	fmt.Println("type one invoke methodTwo...")
}

//当某个类型上增加了某个接口中的所有方法，那么这个类型 和 这个类型的指针，都可以赋值给接口对象；
func ShowOne() {
	var iOne interfaceOne
	var o one = 12
	iOne = o
	iOne.methodOne()
	iOne.methodTwo()

	oReference := &o
	oReference.methodOne()
	oReference.methodTwo()
}

//=================================================================

type two int

func (two *two) methodTwo() {
	fmt.Println("type reference two invoke methodTwo...")
}

func (two *two) methodThree() {
	fmt.Println("type reference two invoke methodThree...")
}

//当某个类型的指针增加了某个接口中的所有方法，那么只有这个类型的指针可以赋值给接口对象；
func ShowTwo() {
	var iTwo interfaceTwo
	var t two = 17

	//iTwo = t

	iTwo = &t
	iTwo.methodTwo()
	iTwo.methodThree()
}
