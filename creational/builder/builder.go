/*
建造者模式
核心解决：	通过多个简单对象一步步构建出复杂的对象
原理:		在抽象建造器中维护着一个目标对象，通过回溯

*/
package builder

import (
	"fmt"
)

///////////////  实例型建造器 //////////////
type Car struct {
	name  string
	brand string
	color string
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

//////////////// 抽象建造器 /////////////
type CarBuilder struct {
	car *Car
}
type Builder interface {
	Brand(string) Builder
	Color(string) Builder
	Name(string) Builder
	Build() *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c CarBuilder) Name(s string) Builder {
	if c.car == nil {
		c.car = &Car{}
	}
	c.car.name = s
	return c
}

func (c CarBuilder) Brand(s string) Builder {
	if c.car == nil {
		c.car = &Car{}
	}
	c.car.brand = s
	return &c
}

func (c CarBuilder) Color(s string) Builder {
	if c.car == nil {
		c.car = &Car{}
	}
	c.car.color = s
	return &c
}

func (c CarBuilder) Build() *Car {
	return c.car
}
