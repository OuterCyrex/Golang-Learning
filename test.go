package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

/*
给Group实现一个Do方法，方法标签如下：
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error)
保证不论Do方法被执行多少次，fn都只被执行一次
注意:key和fn是成对的
这里的不论执行多少次意思是多线程场景下，是如果一瞬间有多个线程访问fn函数，可以保证只执行一次fn操作，共享fn的结果
*/

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	value, ok := g.m[key]
	if !ok {
		newVal, err := fn()
		g.m[key] = &call{wg: sync.WaitGroup{}, val: newVal, err: err}
		g.m[key].wg.Add(1)
		defer g.m[key].wg.Done()
		return newVal, err
	} else {
		value.wg.Wait()
		return value.val, value.err
	}
}

//思路：在多个线程访问Do方法时，每当一个goroutine输入一个Key值，先判断该Key值是否在Group.map[string]内存在，若不存在，则创造一个新的映射
//若已存在该映射，则等待所有该映射key值对应的fn全部运行完，再输出该Key对应fn输出的空接口和err值

func TestFun() (interface{}, error) {
	time.Sleep(1 * time.Second)
	return "result", nil
}
func main() {
	group := new(Group)
	group.m = make(map[string]*call)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			result, err := group.Do("Key", TestFun)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("goroutine %d key: %v\n", i, result)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
