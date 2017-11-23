package abstractfactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	t.Log("创建抽象工厂")
	abstractFactory := new(AbstractFactory)

	t.Log("创建Shape工厂")
	shapeFactory, err := abstractFactory.GetFactory("SHAPE")
	if err != nil {
		t.Log("创建Shape工厂失败")
	}

	t.Log("创建目标实例circle")
	circle, _ := shapeFactory.GetShape("CIRCLE")
	circle.Draw()

	t.Log("创建目标实例rectangle")
	rectangle, _ := shapeFactory.GetShape("RECTANGLE")
	rectangle.Draw()

	t.Log("创建Color工厂")
	colorFactory, err := abstractFactory.GetFactory("COLOR")
	if err != nil {
		t.Log("创建Color工厂失败")
	}

	t.Log("创建目标实例red")
	red, _ := colorFactory.GetColor("RED")
	red.Paint()

	t.Log("创建目标实例yellow")
	yellow, _ := colorFactory.GetColor("YELLOW")
	yellow.Paint()

}
