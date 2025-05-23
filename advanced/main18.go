package main

import (
	"fmt"
	"sync"
)
//线程安全结构体：
type SafeCounter struct {
	mu      sync.Mutex //互斥锁，保护计数器访问
	counter int //需要保护的共享计数器
}

//原子递增方法：
func (s *SafeCounter) Increment() {
	s.mu.Lock() //加锁
	defer s.mu.Unlock()//无论如何 释放锁 使用 defer 确保锁必然释放
	s.counter++ //自增 保证每次递增操作的原子性
}

func main()  {
	//初始化
	var (
		counter = SafeCounter{}
		wg      sync.WaitGroup
		workers = 10
	)
	//并发控制
	wg.Add(workers) //设置需要等待的协程数量
	//启动工作协程：
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()//协程结束时调用 wg.Done()
			//每个协程执行 1000 次递增
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	//等待所有协程完成
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.counter)
}
