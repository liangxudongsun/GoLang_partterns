/*
对象池模式
核心解决：	接口选择问题
原理:		通过方法创建一个对象池，提供外部接口对对象池里的对象进行操作。
			方法能够自动维护对象池。
实践例子:		1、数据库连接池 2、

*/
package objectPool

import (
	"errors"
	"fmt"
	"sync"
)

///////////////  数据结构 //////////////

type Object struct {
	Shape string
	Brand string
	pool  *Pool //对象池
}

type Pool struct {
	//对象池 此处抽象为对象
	objects chan *Object
	//当前活跃对象数
	activeObject int
	//最大对象数
	maxObject int
	//互斥锁
	mu sync.Mutex
}

///////////////  外部接口 //////////////

func NewPool(maxIdle int) *Pool {
	p := Pool{
		objects:      make(chan *Object, maxIdle),
		maxObject:    maxIdle,
		activeObject: 0,
	}
	//初始化池对象
	for i := 0; i < p.maxObject; i++ {
		p.objects <- &Object{
			pool: &p,
		}
	}
	return &p
}

func (p *Pool) Get() (error, *Object) {
	p.mu.Lock()
	defer p.mu.Unlock()

	//当活跃对象数等于最大对象数时候，无法获取到对象
	if p.activeObject == p.maxObject {
		return errors.New("无空闲对象"), nil
	}
	o := <-p.objects
	p.activeObject += 1

	return nil, o
}

func (p *Object) Close() {
	p.pool.mu.Lock()
	defer p.pool.mu.Unlock()

	//返回到对象池
	p.pool.objects <- p
	p.pool.activeObject -= 1
	return
}

func (p *Object) Show() string {
	return fmt.Sprint("Object:Brand=", p.Brand, " Shape=", p.Shape)
}
