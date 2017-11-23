/*
工厂模式
核心解决：	接口选择问题
原理:		通过工厂类中的方法进行目标类的选择。
*/
package factory

import (
	"fmt"
)

///////////////  外部接口 //////////////
type ShapeFactory struct {
}

func (s ShapeFactory) GetShape(shape string) (Shape, error) {
	switch shape {
	case "RECTANGLE":
		return Shape(Rectangle{}), nil
	case "CIRCLE":
		return Shape(Circle{}), nil
	default:
		return nil, nil
	}
}

type Shape interface {
	Draw() //绘图
}

///////////////  实例 //////////////
//方形
type Rectangle struct {
}

func (r Rectangle) Draw() {
	fmt.Println("Rectangle")
}

//圆形
type Circle struct {
}

func (c Circle) Draw() {
	fmt.Println("Circle")
}
