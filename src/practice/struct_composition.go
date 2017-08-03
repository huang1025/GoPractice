package practice

import "fmt"

type Integer int

func (i *Integer) add(num Integer) {
	*i = *i + num
}

type Double int

func (d *Double) add(num Integer) {
	*d = *d + 1
}

//=================================================================

type person struct {
	age Integer
}

//如果成员有名称，那么可以直接使用成员，调用成员中增加的方法；
func RunPerson() {
	p := person{}
	fmt.Println(p.age) //0

	p.age.add(12)
	fmt.Println(p.age) //12

	p.age.add(p.age)
	fmt.Println(p.age) //24
}

//=================================================================

type personTwo struct {
	Integer
}

//如果成员没有名称，就属于匿名组合，结构体可以直接调用匿名成员上增加的方法；
func RunPersonTwo() {
	p := personTwo{}
	p.add(12)
	fmt.Println(p) //{12}
}

//=================================================================

type personThree struct {
	Integer
	Double
}

//如果成员拥有相同的方法，那么即使是匿名成员，也不能直接调用，需要借用匿名成员的类型，进行下一步的调用；
func RunPersonThree() {
	p := personThree{}
	p.Integer.add(12)
	fmt.Println(p) //{12 0}
}

//=================================================================

type personFour struct {
	Double
}

//如果匿名成员和该类型都拥有同一个方法，那么直接调用的将是该类型的方法；
func (four personFour) add(num Integer) {
	fmt.Println("personFour type is invoking add method...")
}

func RunPersonFour() {
	four := personFour{}
	four.add(12)
}
