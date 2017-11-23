/*
抽象工厂模式
核心解决：	接口选择问题
原理:		创建抽象类，通过抽象类中的方法进行工厂类的选择，
			通过工厂类的方法最后进行目标类的选择。

*/
package abstractfactory

import (
	"errors"
	"fmt"
)

///////////////  外部接口 //////////////
type AbstractFactory struct {
}

func (a AbstractFactory) GetFactory(factory string) (Factory, error) {
	switch factory {
	case "COLOR":
		return Factory(new(ColorFactory)), nil
	case "SHAPE":
		return Factory(new(ShapeFactory)), nil
	default:
		return nil, errors.New("Unsupported factory")
	}

}

type Factory interface {
	GetColor(color string) (Color, error)
	GetShape(shape string) (Shape, error)
}

////////////////////////// 工厂1 ///////////////////////////
type ShapeFactory struct {
}

func (s *ShapeFactory) GetColor(color string) (Color, error) {
	return nil, nil
}
func (s *ShapeFactory) GetShape(shape string) (Shape, error) {
	switch shape {
	case "RECTANGLE":
		return Shape(new(Rectangle)), nil
	case "CIRCLE":
		return Shape(new(Circle)), nil
	default:
		return nil, errors.New("Unsupport shape")
	}
}

type Shape interface {
	Draw() //绘图
}

///////////////  实例 //////////////
//方形
type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Rectangle")
}

//圆形
type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Circle")
}

////////////////////////// 工厂2 ///////////////////////////
type ColorFactory struct {
}

func (s *ColorFactory) GetShape(Shape string) (Shape, error) {
	return nil, errors.New("Invalid mode")
}

func (s *ColorFactory) GetColor(color string) (Color, error) {
	switch color {
	case "RED":
		return Color(new(Red)), nil
	case "YELLOW":
		return Color(new(Yellow)), nil
	default:
		return nil, errors.New("Unsupport color")
	}
}

type Color interface {
	Paint() //绘图
}

///////////////  实例 //////////////
//方形
type Red struct {
}

func (r *Red) Paint() {
	fmt.Println("Red")
}

//圆形
type Yellow struct {
}

func (y *Yellow) Paint() {
	fmt.Println("Yellow")
}
