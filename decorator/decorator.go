/*
装饰器模式
核心解决：	向一个现有的对象添加新的功能，同时又不改变其结构。
原理:		在JAVA中通过新建一个类继承原有的类，并在类中实现原有方法基础上增加新功能。
			在Go中可以通过将原有方法作为参数传入，并在新的装饰器中实现。
*/
package decorator

import (
	"fmt"
)

///////////////  目标方法   //////////////////////
type TargetMethod func() string

//目标对象
type TargetObject struct {
	Color string
}

//目标方法
func (t TargetObject) Draw() string {
	return "The color is " + t.Color
}

//目标公共方法
func Draw() string {
	return "Public draw"
}

//////////////   修饰器   //////////////////////
//目标方法装饰器
func DrawDecorate(d TargetMethod) TargetMethod {
	r := func() string {
		fmt.Println("Decorate begin")
		//运行主方法
		k := d()
		fmt.Println("Decorate end")
		return k
	}
	return r
}

//目标对象装饰器
func DrawObjectDecorate(d *TargetObject) {
	fmt.Println("Decorate begin")
	//运行主方法
	fmt.Println(d.Draw())
	fmt.Println("Decorate end")
}
