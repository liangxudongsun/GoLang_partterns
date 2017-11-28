/*
cancel模式
核心解决：	控制协程退出
原理:		通过向协程参数中加入context，以实现对子协程的控制。
			控制方式: 1、父协程调用cancel() 2、超时自动退出
场景:		RPC,HTTP,需要对子协程进行控制的情况
*/
package cancel

import (
	"context"
	"fmt"
	"time"
)

//1、父级Context被关闭时候子级关闭
func CancelPattern() {
	//伪
	k := func(ctx context.Context, id int) {
		defer fmt.Println(id, " has exited")

		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(id, " in goroutince")
			}
		}
	}

	parCtx := context.Background()
	ctx, cancel := context.WithCancel(parCtx)

	//TestCode
	for i := 1; i != 10; i++ {
		go k(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

//2、父级ctx被关闭或者达到超时条件时候子级关闭
//注意:这里的超时时间是指从每个函数调用入口开始的时候算起的时间
//	   若需要指定时间，用context.WithDeadline()
//	   WithTimeout 等价于 WithDealline(ctx,time.Now().Add(Timeout))
func CancelWithTimeoutParttern() {
	//伪
	k := func(ctx context.Context, id int) {
		defer log.Println(id, " has exited")

		for {
			select {
			case <-ctx.Done():
				return
			default:
				log.Println(id, " in goroutince")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//TestCode
	for i := 1; i != 10; i++ {
		go k(ctx, i)
	}

	//Wait for timeout
	time.Sleep(10 * time.Second)

}
