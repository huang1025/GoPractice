package practice

import "fmt"

type IntegerOne int

func (i IntegerOne) add() {
	i = i + 1
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
	*i += 1
}

//当给某个类型的指针类型增加方法后，这个类型 和 这个类型的指针 都可以使用这个增加的方法；
func ShowIntegerTwo() {
	var two IntegerTwo = 17
	two.add()

	twoReference := &two
	twoReference.add()
}

//=================================================================

//当给类型增加方法后，不管是类型 还是类型的指针类型 调用方法都不能更改原有的值；
//当给类型的指针类型增加方法后，不管是 类型 还是 类型的指针类型，调用方法都能改变原有的值；
func RunIntegerOne() {
	var i IntegerOne = 12
	i.add()
	fmt.Println(i) //12

	iReference := &i
	iReference.add()
	fmt.Println(*iReference) //12

	var i2 IntegerTwo = 12
	i2.add()
	fmt.Println(i2) //13

	i2Renfence := &i2
	i2Renfence.add()
	fmt.Println(*i2Renfence) //14
}

//=================================================================

type IntegerThree int

func (i IntegerThree) add() {}

// method redeclared: IntegerThree.add
//func (i *IntegerThree)add(){}