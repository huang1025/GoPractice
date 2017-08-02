package practice

import "fmt"

type IntegerOne int

func (i IntegerOne) add() {
	fmt.Println("IntegerOne invoke add...")
}

//当给某个类型增加方法后，这个类型和这个类型的引用都可以使用增加的方法；
func ShowIntegerOne() {
	var one IntegerOne = 12
	one.add()

	oneReference := &one
	oneReference.add()
}

//=================================================================

type IntegerTwo int

func (i *IntegerTwo) add() {
	fmt.Println("IntegerTwo invoke add...")
}

//当给某个类型的指针类型增加方法后，这个类型 和 这个类型的指针 都可以使用这个增加的方法；
func ShowIntegerTwo() {
	var two IntegerTwo = 17
	two.add()

	twoReference := &two
	twoReference.add()
}
