package serve

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Listen(addr string) {

	server := &http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Web server listen on " + addr)

	// 以协程方式启用监听，防止阻塞后续的系统信号处理
	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("%s\n", err)
		}
	}()

	// 等待中断信号正常关闭服务器
	quit := make(chan os.Signal)
	// `kill` 默认发送 SIGTERM 信号
	// `kill -2` 发送 SIGINT 信号
	// `kill -9` 发送 SIGKILL 信号，无法捕获，不处理
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// 获取信号，如果没有则保持阻塞
	<-quit

	log.Println("Server shutting down...")

	// 通知服务器还有5秒的时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器而不中断任何活动连接
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
