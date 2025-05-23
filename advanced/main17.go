package main

import (
	"fmt"
	"sync"
)

var wg17 sync.WaitGroup

//生产中
func producer(intChan chan int)  {
	defer wg17.Done()
	for i := 1; i <=100 ; i++ {
		intChan<-i
		fmt.Println("写入...:", i)
	}
	close(intChan)
}

//消费者
func customer(intChan chan int)  {
	defer wg17.Done()
	for v := range intChan{
		fmt.Println("读取....：",v)
	}
}

func main()  {
	fmt.Println("开始....")
	intChan := make(chan int, 100)
	wg17.Add(2)
	go producer(intChan)
	go customer(intChan)
	fmt.Println("结束....")
	wg17.Wait()
}
