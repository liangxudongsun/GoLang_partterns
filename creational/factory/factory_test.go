package factory

import (
	"testing"
)

func TestFactory(t *testing.T) {
	//创建公共实例
	t.Log("创建公共实例")
	shapeFactory := new(ShapeFactory)
	//创建目标实例
	t.Log("创建目标实例circle")
	circle, _ := shapeFactory.GetShape("CIRCLE")
	circle.Draw()

	t.Log("创建目标实例rectangle")
	rectangle, _ := shapeFactory.GetShape("RECTANGLE")
	rectangle.Draw()
}
