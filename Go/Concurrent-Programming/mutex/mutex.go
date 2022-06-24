package main

import (
	"fmt"
	"sync"
)

// Example01 不加锁会产生什么情况 对一个数开启10个协程 加100  期望结果 1000
func Example01() {
	count := 0
	//使用sync.WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

// Example02 引入Mutex 解决data race
func Example02() {
	var mu sync.Mutex // 互斥锁保护计数器，sync.Mutex 的零值是还没有 goroutine等待的未加锁状态

	count := 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("count: %d\n", count)
}

func main() {
	//Example01()
	//Example02()
	//Example03()
	Example04()
}

// Counter01 Counter sync.Mutex 的多种使用方式
type Counter01 struct {
	mu    sync.Mutex
	Count uint64
}

type Counter02 struct {
	sync.Mutex // 嵌套
	Count      uint64
}

// Example03 在struct调用 Mutex
func Example03() {
	var Counter02 = Counter02{}
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				Counter02.Lock()
				Counter02.Count++
				Counter02.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter02.Count: ", Counter02.Count)
}

type Counter03 struct {
	CounterType int
	name        string

	sync.Mutex
	count uint64
}

// Incr +1的方法内部使用互斥锁保护
func (c *Counter03) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

// Example04 封装加锁的方法 不对外暴露细节
func Example04() {
	var Counter03 = Counter03{}
	var wg sync.WaitGroup

	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				Counter03.Incr()
			}
		}()
	}
	for k := 0; k < 10; k++ {
		go func() {
			defer wg.Done()
			Counter03.ReadCount()
			//fmt.Println("Counter03.Count: ", Counter03.count)
			//fmt.Println("Counter03.Count(): ", Counter03.ReadCount())
		}()
	}

	wg.Wait()
	//fmt.Println("-Counter03.Count: ", Counter03.count)
	//fmt.Println("-Counter03.Count(): ", Counter03.ReadCount())

	fmt.Println("Counter03.Count: ", Counter03.count)
	fmt.Println("Counter03.Count(): ", Counter03.ReadCount()) // 因为这里有调用了一次
}

// ReadCount 读取count的方法 为什么还要有读加锁的方法，通过前面的案例 确实可以将数据累加成功
func (c *Counter03) ReadCount() uint64 {
	c.Lock()
	defer c.Unlock()

	c.count++
	return c.count
}
