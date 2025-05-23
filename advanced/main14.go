package main

import "fmt"

//定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法
type  Shape interface {
	Area()
	Perimeter()
}
//创建 Rectangle
type Rectangle struct {

}
//Circle 结构体
type Circle struct {

}

func (rectangle Rectangle) Area()  {
	fmt.Println("Rectangle....Area.....")
}
func (rectangle Rectangle) Perimeter()  {
	fmt.Println("Rectangle....Perimeter.....")
}

func (circle Circle) Area()  {
	fmt.Println("Circle....circle.....")
}
func (circle Circle) Perimeter()  {
	fmt.Println("Circle....circle.....")
}

func main()  {

	r := Rectangle{}
	r.Area()
	r.Perimeter()
	c := Circle{}
	c.Area()
	c.Perimeter()

}