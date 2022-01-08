package main

import "fmt"

/*
定义结构体 employee，
它有一个 salary 字段，
给这个结构体定义一个方法 giveRaise
来按照指定的百分比增加薪水。
*/

type employee struct {
	salary float32
}

func main() {
	E := new(employee)
	E.salary = 8000
	afterTheSalaryIncrease := E.giveRaise(40.0)
	fmt.Printf("在跳槽后可以拿到%v的工资\n", afterTheSalaryIncrease)
}

func (Eee *employee) giveRaise(percent float32) (salary float32) {
	salary = Eee.salary * (1.0 + percent/100.0)
	return
}
