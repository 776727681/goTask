package main

import "fmt"

func main()  {
	num := 1
	fmt.Println("修改前:",num)
	task(&num)
	fmt.Println("修改后:",num)
}

func task(ptr *int)  {
	// 解引用指针并增加10
	*ptr += 10
}

