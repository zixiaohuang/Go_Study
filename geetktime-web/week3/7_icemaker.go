package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	//id int // 类型不一样，交错执行有可能直接panic
	name string
}

// interface 内部有两个machine word，所以会出现交错的情况，Ben says, "Hello my name is Jerry"
func (b *Ben)Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

//type Jerry struct {
//	name string
//}

type Jerry struct { // string类型底层刚好就这样描述表示的，所以刚好不报错
	field1 *[5]byte
	field2 int
}

func (j *Jerry)Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", string(j.field1[:]))
	//fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.name)
}

func main() {
	//var ben = &Ben{id: 10, name: "Ben"}
	var ben = &Ben{"Ben"}
	var jerry = &Jerry{"Jerry"}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()
	for {
		maker.Hello()
	}
}