// go build -race lesson3_detecingrace.go
// 发现detecing race，整个进程stop掉
// 会告诉你，哪个goroutine、哪个内存地址产生了读行为 Previous read at 0x000001201788 by goroutine 8:
// 哪个地方产生了写行为Write at 0x000001201788 by goroutine 7:
package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0
func main()  {
	for routine:= 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine) // 开了很多个goroutine，同时产生读写，就会发现data race
	}
	Wait.Wait()
	fmt.Printf("Final Counter %d\n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter // 读行为
		time.Sleep(1* time.Nanosecond)
		value++
		Counter = value // 赋值 写行为
	}
	Wait.Done()
}