// +build !windows

package quit

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func safeExit() {
	// 创建监听退出chan
	s := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for c := range s {
			switch c {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", c)
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1", c)
			case syscall.SIGUSR2:
				fmt.Println("usr2", c)
			default:
				fmt.Println("other", c)
			}
		}
	}()
}
