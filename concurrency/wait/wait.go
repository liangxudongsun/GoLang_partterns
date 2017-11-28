/*
wait模式
核心解决：	等待协程退出
原理:		引入sync中的waitgroup功能，在主线程中对每个协程都进行计数
			主线程或者父线程可控制等待所有协程退出后再进行其他操作。
场景:		需要确保所有子协程退出才能进行操作
*/
package wait

import (
	"fmt"
	"sync"
	"time"
)

func WaitPattern() {
	//伪
	var wg sync.WaitGroup

	k := func(id int) {
		defer fmt.Println(id, " has exited")
		defer wg.Done()

		fmt.Println(id, " in goroutince")
		time.Sleep(2 * time.Second)
		return
	}

	for i := 1; i != 10; i++ {
		go k(i)
	}

	//等待所有协程退出
	wg.Wait()
	fmt.Println("所有协程已退出")

}
