// week5：参考 Hystrix 实现一个滑动窗口计数器

package main

import (
	"container/ring"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// CounterBucket 采样桶：统计当前时间windowStart单位s内，请求总数total和请求成功数success
type CounterBucket struct {
	success int64
	total int64
	windowStart int64
}

func (c *CounterBucket) AddRequest(isSuccess bool) {
	atomic.AddInt64(&c.total, 1)
	if isSuccess{
		atomic.AddInt64(&c.success, 1)
	}
}

func (c *CounterBucket) GetBucketCounter() (total, success int64) {
	return c.total, c.success
}

func (c *CounterBucket) Reset(windowStart int64) {
	atomic.StoreInt64(&c.total, 0)
	atomic.StoreInt64(&c.success, 0)
	atomic.StoreInt64(&c.windowStart, windowStart)
}

type RollingWindowCounter struct {
	r *ring.Ring
	head *ring.Ring
	tail *ring.Ring
}

func (rw *RollingWindowCounter) Init(size int) {
	rw.r = ring.New(size + 1)
	rw.head = rw.r
	rw.tail = rw.r

	for i:= 0; i < rw.r.Len(); i++ {
		rw.head.Value = &CounterBucket{
			success: 0,
			total: 0,
			//windowStart: time.Now().UnixNano(),
			windowStart: time.Now().Unix(), // s=1000ms
		}
		//temp := rw.head.Value
		rw.head = rw.head.Next()
	}
	rw.head = rw.tail.Next()
	//rw.r.Do(fmt.Printf(""))

	rw.r.Do(func(p interface{}) {
		fmt.Printf("%+v\n", p)
	})
}

func (rw *RollingWindowCounter) AddCount(isSuccess bool) error {
	if rw.r == nil {
		return errors.New("RollingWindowCounter not initial err")
	}

	cb, ok := rw.head.Value.(*CounterBucket)
	if !ok {
		return errors.New("RollingWindowCounter type err")
	}
	timeNow := time.Now().Unix()
	//timeNow := time.Now().UnixNano()
	//fmt.Printf("timeNow = %d\n",timeNow)

	// 在相同计算单位下，直接调用
	if timeNow == cb.windowStart{
		cb.AddRequest(isSuccess)
		return nil
	}

	// 将超过了计算单位前面的桶重置，
	oldWindowStart := cb.windowStart
	diff := timeNow - oldWindowStart
	fmt.Printf("diff = %d\n", diff)
	for i:= int64(1); i <=diff; i++ {
		rw.head =rw.head.Next()
		if rw.head == rw.tail {
			rw.tail = rw.tail.Next()
		}

		cb2, ok := rw.head.Value.(*CounterBucket)
		if !ok {
			return errors.New("RollingWindowCounter type err")
		}
		cb2.Reset(oldWindowStart + i)
		if cb2.windowStart == timeNow {
			//print("cb2 AddRequest\n")
			cb2.AddRequest(isSuccess)
			return nil
		}
	}
	return nil
}

func (rw*RollingWindowCounter) Sum()(total, success int64) {
	total, success = 0, 0
	for r := rw.tail; r != rw.head; r = r.Next() {
		if cb, ok := r.Value.(*CounterBucket); ok {
			//fmt.Printf("counterbucket add = %d, r add = %d\n", &cb, &r)
			total += cb.total
			success += cb.success
		}
	}
	return
}

func main() {
	rollingWindowCounter := &RollingWindowCounter{}
	//
	rollingWindowCounter.Init(10) // 10个 CounterBucket
	//fmt.Printf("unix nano = %d, cur time = %d\n", time.Now().UnixNano(), time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wg := sync.WaitGroup{}
	fmt.Println(time.Now())
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i:= 0; i < 4000; i++ {
				random := r.Intn(10)
				//fmt.Printf("random %d\n",random)
				isSuccess := false
				if random < 8 {
					isSuccess = true
				}
				_ = rollingWindowCounter.AddCount(isSuccess)
				//time.Sleep(time.Millisecond * time.Duration(random))
				time.Sleep(time.Millisecond * time.Duration(random))
			}
		}()
	}
	wg.Wait()
	//fmt.Println(time.Now())
	total,success := rollingWindowCounter.Sum()
	fmt.Printf("total: %d, success: %d\n", total, success)
}

