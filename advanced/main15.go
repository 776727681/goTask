package main

import "fmt"

type Person struct {
	Name string
	Age int
}
type Employee struct {
	Person // 匿名嵌入Person结构体
	EmployeeID int
}

func (e Employee) PrintInfo()  {
	fmt.Println(e.Age)
	fmt.Println(e.Name)
	fmt.Println(e.EmployeeID)
}


func main()  {
	e := Employee{
		Person:Person{
			Name: "张三",
			Age: 20,
		},
		EmployeeID: 100,
	}
	e.PrintInfo()
}
