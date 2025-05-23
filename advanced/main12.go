package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

func one()  {
	defer wg.Done()
	for i:=1; i<=10;i++ {
		if i%2 > 0{
			fmt.Println("奇数：",i)
		}
	}
}

func two()  {
	defer wg.Done()
	for i:=1; i<=10;i++ {
		if i%2 == 0{
			fmt.Println("偶数：",i)
		}
	}
}


func main() {
	fmt.Println("start....")
	wg.Add(2)
	go one()
	go two()
	wg.Wait()
	fmt.Println("end......")
}
