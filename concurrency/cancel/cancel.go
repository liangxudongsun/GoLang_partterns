/*
cancel模式
核心解决：	控制协程退出
原理:		通过向协程参数中加入context，以实现对子协程的控制。
*/
package cancel

import (
	"context"
	"fmt"
	"time"
)

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

	for i := 1; i != 10; i++ {
		go k(ctx, i)
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
