package main

import (
	"fmt"
	"sync"
)

var  wg16 sync.WaitGroup//等待一组并发操作完成
var lock sync.RWMutex //读写锁
func read (intChan chan int) {
	defer wg16.Done()
	for v := range intChan {
		lock.RLock()
		fmt.Println("read..:",v)
		lock.RUnlock()
	}
	
}
func write(intChan chan int)  {
	defer wg16.Done()
	for i:=1; i<=10; i++ {
		lock.Lock()
		intChan <- i
		lock.Unlock()
		fmt.Println("write...:", i)
	}
	close(intChan)

}
func main()  {
	fmt.Println("start....")
	intChan := make(chan int, 10)
	wg16.Add(2)
	go read(intChan)
	go write(intChan)
	wg16.Wait()
	fmt.Println("end...")


}