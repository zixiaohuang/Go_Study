package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m_lock *sync.RWMutex
	once sync.Once
	output string = "./memory_info.txt"
	NUM = 10
	write_size = [6]int{1, 10, 20, 30, 40, 50} //对应val多少个byte
)

func WriteInfo(content string) {
	m_lock.Lock()
	fp, err := os.OpenFile(output, os.O_CREATE| os.O_APPEND|os.O_RDWR, 0666)
	defer fp.Close()
	if _, err = fp.WriteString(content); err != nil {
		fmt.Println(err)
	}
	m_lock.Unlock()
}

func main() {
	m_lock = new(sync.RWMutex)
	//c, err := redis.Dial("tcp", "8.134.209.208:6379")
	pool := &redis.Pool{
		MaxIdle: 16,
		MaxActive: 1024,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "8.134.209.208:6379")
		},
	}
	//if err != nil {
	//	fmt.Println("conn redis failed, err:", err)
	//	return
	//}
	//defer c.Close()
	//once_print := func() {
	//	outstr := fmt.Sprintf("memory info before set:")
	//	res, err := redis.String(c.Do("info", "memory"))
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	outstr += res
	//	WriteInfo(outstr)
	//}

	wg := sync.WaitGroup{}
	for i:= 0; i < 6; i++ {
		wg.Add(1)

		go func(size int) {
			defer wg.Done()
			//once.Do(once_print) // 设置kv前的info 信息，只执行一次
			c := pool.Get()
			defer c.Close()
			val := ""
			for j := 0; j < size; j++ {
				val += "h" // 每个字符一个byte
			}
			outstr := fmt.Sprintf("memory info of %d byte:", size)
			outstr += "before info: \n"
			res, err := redis.String(c.Do("info", "memory"))
			if err != nil {
				fmt.Println(err)
				return
			}
			outstr += res
			//kv := make(map[string]interface{})
			offset := time.Now().UnixNano()
			for i := 0; i < NUM; i++ { // 插入多少次
				k := "test_" + strconv.FormatInt(offset, 10) // 测试一下
				_, err = c.Do("Set", k, val)
				if err != nil {
					fmt.Println(err)
				}
				offset++
			}
			//_, err = c.Do("MSet", kv)
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
			outstr += "after info: \n"
			res, err = redis.String(c.Do("info", "memory"))
			if err != nil {
				fmt.Println(err)
				return
			}
			outstr += res
			outstr += "\n"
			WriteInfo(outstr)
		}(write_size[i])
	}
	wg.Wait()
	c := pool.Get()
	_, err := c.Do("flushall") //测试完清空实例
	if err != nil {
		fmt.Println(err)
	}
}