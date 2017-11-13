/*
原型模式
核心解决：	接口选择问题
原理:		实现了一个原型接口，该接口用于创建当前对象的克隆。
实践例子:		1、数据对象的缓存

*/
package prototype

///////////////  数据结构 //////////////
//原型库
type ShapeCache struct {
	shapeList map[string]Shape
}

func NewShapeCache() *ShapeCache {
	return &ShapeCache{
		shapeList: make(map[string]Shape, 0),
	}
}

func (s ShapeCache) LoadCache() {
	//实例化所有原型对象
	s.shapeList["Circle"] = Shape(Circle{
		shape: "Circle",
		brand: "Audi",
	})

	s.shapeList["Rectangle"] = Shape(Rectangle{
		shape: "Rectangle",
		brand: "HONDA",
	})
}

func (s ShapeCache) GetShape(shape string) (error, Shape) {
	return nil, s.shapeList[shape]
}

//原型模板
type Shape interface {
	GetType() string
}

//原型对象1
type Circle struct {
	shape string
	brand string
}

func (c Circle) GetType() string {
	return c.shape
}

//原型对象2
type Rectangle struct {
	shape string
	brand string
}

func (c Rectangle) GetType() string {
	return c.shape
}
