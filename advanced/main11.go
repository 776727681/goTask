package main

import "fmt"

func main()  {
	nums := []int{1,2,3,4,5}
	fmt.Println("修改前：",nums )
	task11(&nums)
	fmt.Println("修改后：",nums )
}

func task11(slicePtr *[]int)  {
	for i:=range *slicePtr{
		(*slicePtr)[i] = (*slicePtr)[i] * 2
	}
}
