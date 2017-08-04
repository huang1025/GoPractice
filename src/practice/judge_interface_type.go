package practice

import "fmt"

type Animal interface {
}

type Person struct {
	name string
	age  int
}

func RunInterfaceTypeJudge() {
	var a Animal = Person{"huang", 17}
	one, two := a.(int)
	fmt.Println(one) //0
	fmt.Println(two) //false

	three, four := a.(Person)
	fmt.Println(three) //{huang 17}
	fmt.Println(four)  //true

	five, six := a.(*Person)
	fmt.Println(five) //<nil>
	fmt.Println(six)  //false

	var b Animal = &Person{"huang", 17}
	seven, eight := b.(*Person)
	fmt.Println(seven)      //&{huang 17}
	fmt.Println(seven.name) //huang
	fmt.Println(eight)      //true

	//var b Person = Person{"huang", 17}
	//b.(Person)           //invalid type assertion
}
