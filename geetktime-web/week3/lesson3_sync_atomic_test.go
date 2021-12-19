package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

//type Config struct {
//	a []int
//}

func (c *Config) T() {}

//func BenchmarkFirstVersion(b *testing.B) {
//	cfg := &Config{}
//
//	go func() {
//		i := 0
//		for {
//			i++
//			cfg.a = []int{i, i+1, i + 2, i + 3, i + 4, i + 5};
//		}
//	}()
//
//	var wg sync.WaitGroup
//	for n := 0; n < 4; n++ {
//		wg.Add(1)
//		go func() {
//			for n:=0;n<100;n++{
//				fmt.Printf("%v\n", cfg)
//			}
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}

func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value // 读特别特别多，可以用atomic.Value，性能可能比读写锁高
	v.Store(&Config{}) // 老对象

	go func() {
		i := 0
		for {
			i++
			cfg := &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}} //赋值操作极少，完全可以用新对象替换
			v.Store(cfg) // 每次new一个存起来
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				cfg := v.Load().(*Config)// Load 出来， 断言成Config对象
				cfg.T() // 打印出来
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
}

func BenchmarkMutex(b *testing.B) {
	var l sync.RWMutex
	var cfg *Config

	go func() {
		i := 0
		for {
			i++
			l.Lock() // 赋值再解锁
			cfg = &Config{a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}}
			l.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < b.N; n++ {
				l.RLock()
				cfg.T() // 打印出来
				l.RUnlock()
			}
			wg.Done()
		}()
	}
}
