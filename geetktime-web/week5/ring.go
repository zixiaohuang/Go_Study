package main

import (
	"container/ring"
	"fmt"
)

func main()  {
	r := ring.New(3)
	//n := r.Len()
	for i :=1; i <= 3; i++ {
		r.Value = i
		fmt.Printf("i = %d, %+v\n", i, r)
		r = r.Next()
		//fmt.Printf("init head address = %d\n", r) //这样打印是不行的， r相当于是同一个指针，只是指向不同的Ring结构
	}

	//for j := 0; j < n; j++ {
	//	fmt.Println(r.Value)
	//	r = r.Next()
	//}
	r.Do(func(p interface{}) {
		fmt.Printf("%+v\n", p)
	})
}
