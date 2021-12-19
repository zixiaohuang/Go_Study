package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int

	// 三个人并行执行
	// 调用广告
	g.Go(func() error {
		a = xxx
		return errors.New("test")
	})

	// 调用AI
	g.Go(func() error {
		b = xxx
	})

	// 调用运营平台
	g.Go(func() error {
		c = xxx
	})

	err:= g.Wait()
	// 在wait方法后对abc做数据组装
	fmt.Printf(err)

	fmt.Println(ctx.Err())
}
