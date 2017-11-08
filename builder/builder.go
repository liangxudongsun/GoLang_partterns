/*
建造者模式
核心解决：	通过多个简单对象一步步构建出复杂的对象
原理:		经常被全局使用的类，在内部自行创建实例，并且保证只有一次创建，
			对外提供访问这个唯一对象的接口，可以直接访问，无需再创建。
			在Go中，通过sync包中once功能保证创建实例的方法只会被调用一次
*/
package builder

import (
	"fmt"
)

type Builder interface {
	Brand(string) Builder
	Color(string) Builder
	Name(string) Builder
	Show()
}

///////////////  实例 //////////////
type Car struct {
	name  string
	brand string
	color string
}

///////////////  外部接口 //////////////

func NewCar() Builder {
	c := new(Car)
	return c
}

func (c Car) Name(s string) Builder {
	c.name = s
	return c
}

func (c Car) Brand(s string) Builder {
	c.brand = s
	return &c
}

func (c Car) Color(s string) Builder {
	c.color = s
	return &c
}

func (c Car) Build() *Car {
	return &c
}

func (c Car) Show() {
	fmt.Println("Car: name=", c.name,
		" brand=", c.brand,
		" color=", c.color)
}
