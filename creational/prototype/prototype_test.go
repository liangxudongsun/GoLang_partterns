package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	t.Log("创建原型库")
	cache := NewShapeCache()

	t.Log("加载原型")
	cache.LoadCache()

	t.Log("获取原型Circle")
	_, circle := cache.GetShape("Circle")
	t.Log("Circle.Shape=", circle.GetType())

	t.Log("获取原型Rectangle")
	_, rectangle := cache.GetShape("Rectangle")
	t.Log("Rectangle.Shape=", rectangle.GetType())
}
