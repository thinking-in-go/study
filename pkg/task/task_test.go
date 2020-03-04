package task

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

type PrintJob struct {
	Id int
}

func (p *PrintJob) Do() {
	fmt.Printf("n=%d\n", p.Id)
}

// 接收退出信号
func registerSignal(cancel context.CancelFunc) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, syscall.SIGKILL)
	for s := range ch {
		fmt.Printf("receive signal :%v\n", s)
		cancel()
	}
}

func TestTask_Push(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)

	done := make(chan struct{})
	go registerSignal(cancel)

	task := New(ctx, done, 10)
	var n int32 = 1

	// 模拟不停的发送数据
	go func() {
		for {
			b := task.Push(&PrintJob{Id: int(n)})
			if b {
				atomic.AddInt32(&n, 1)
			} else {
				return
			}
			time.Sleep(time.Millisecond * 200)
		}
	}()

	<-done
	// 打印一共发送了多少条任务
	fmt.Printf("total send Id=%d\n", atomic.LoadInt32(&n)-1)
}
