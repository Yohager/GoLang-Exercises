package main

import "fmt"

type Employee struct {
	name     string
	currency string
	salary   int
}

//写一个方法
func (e Employee) displaySalary() {
	fmt.Printf("name: %s, salary: %s %d\n", e.name, e.currency, e.salary)
}

//写一个函数
func displaySalary2(e Employee) {
	fmt.Printf("name: %s, salary: %s %d\n", e.name, e.currency, e.salary)
}

func main() {
	emp1 := Employee{
		name:     "Alice",
		currency: "$",
		salary:   1000,
	}

	emp1.displaySalary()
	displaySalary2(emp1)
}
