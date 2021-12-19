package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	tr := NewTracker()
	go tr.Run() // 让调用者决定使用goroutine，让调用者自己管理生命周期
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 *time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}

func NewTracker() *Tracker {
	return &Tracker{
		// 引入一个channel，worker工作模型，只需要启动两三个goroutine去后台消费channel里面的数据
		// 不需要启动大量的goroutine。使用channel来做数据的buffer
		ch : make(chan string, 10),
	}
}

// Tracker knows how to track events for the application
type Tracker struct {
	ch chan string
	stop chan struct{} //需要一种信号让channel它暂停
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
		case t.ch <- data:
			return nil
		case <-ctx.Done():
			return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch { // 从goroutine里面不断消费数据，然后上报
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t * Tracker)Shutdown(ctx context.Context) {
	// 调用shutdown：1。先阻断数据的发送，阻断数据发送后Run函数退出 2。收到stop信号，stop信号一收到后我们shutdown就能退出

	// 写的owner才能决定channel的生命周期
	// 做不到，有一个不建议的做法，往channel里面发送一个nil，这样会导致一些事件的丢失
	//先让http shutdown，再调用tracker的shutdown，所有写的人都停止了，才能调用。因为channel关闭了还有人写会导致panic
	close(t.ch) //暂停，不会发送数据
	select {
	case <-t.stop: // 通过stop知道什么时候退出
	case <-ctx.Done(): // shutdown退出时间可能会很长，做一个超时
	}
}

