// 从url拖一些数据出来
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type Resource struct {
	url string
	polling bool // 是否正在爬取
	lastPolled int64 // 上次爬取资源的时间
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex // 互斥锁
}

// 互斥锁方式，很啰嗦
func Poller(res *Resources) {
	// 可以被很多歌goroutine同时访问使用
	for {
		// get the least recently-polled Resource
		// and mark it as being polled
		res.lock.Lock() // 使用锁保证happen before
		var r *Resource
		// 找资源、过滤
		for _, v:= range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil { // 正在爬，把状态设置一下
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil { // 没找到资源，跳出循环
			continue
		}
		// poll the URL
		// do something

		// update the Resource's polling and lastPolled
		res.lock.Lock()
		r.polling = false
		nano, _:= time.ParseDuration("8s")
		r.lastPolled = nano.Nanoseconds()
		res.lock.Unlock()
	}
}


// channel方式，chan内部也是用互斥锁来实现。通过这个包装，用Go原生提供的基础类型去做，写代码就非常简单
type Resource string

func Poller(in, out chan *Resource) {
	for r:= range in { // 从in拿到url
		// poll the URL

		// send the processed Resource to out
		out <- r // 爬取完后再发到out
	}
}