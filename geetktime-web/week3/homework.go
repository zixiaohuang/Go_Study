// 基于errgroup 实现一个http server的启动和关闭，以及linux signal信号的注册和处理，要保证能够一个退出，全部注销退出
package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serveApp() error{
	// DefaultServeMutex是一个全局变量，所有代码都可以修改它，注册处理器可能会冲突
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "server Start")
	})
	if err := http.ListenAndServe(":8080", mux); err!=nil{
		return err
	}
	return nil
}

func serveDebug() error {
	mux := http.NewServeMux()
	return http.ListenAndServe("8081",mux)
}

func serveSignalListen() error {
	signalChannel := make(chan os.Signal, 1)

	// Notify函数让signal包将输入信号转发到c。如果没有列出要传递的信号，会将所有输入信号传递到c；否则只传递列出的输入信号。
	//
	//signal包不会为了向c发送信息而阻塞（就是说如果发送时c阻塞了，signal包会直接放弃）：
	//调用者应该保证c有足够的缓存空间可以跟上期望的信号频率。对使用单一信号用于通知的通道，缓存为1就足够了。
	//
	// SIGINT control-C SIGTERM 软件终止
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	// select case没收到信号前会阻塞
	select {
	case sig := <-signalChannel:
		return errors.Errorf("Get quit signal: %v", sig)
	}
}

func main() {
	g, ctx:= errgroup.WithContext(context.Background())

	g.Go(serveApp)
	g.Go(serveDebug)
	g.Go(serveSignalListen)
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err:= g.Wait(); err != nil {
		log.Printf("exist err: %+v", err)
	}
}