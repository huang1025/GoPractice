package main

import "fmt";

func main() {
	var w Work
	w = new(Employee)
	w.work("文档工作")
	w.sleep("1")

}

type Work interface {
	work(workName string)
	sleep(time string)
}

type Employee struct {
	name   string
	age    int
	height int
}

func (emp Employee) work(workName string) {
	fmt.Print("正在做" + workName)
}

func (emp Employee) sleep(time string) {
	fmt.Print("工人休息" + time + "小时;")
}
