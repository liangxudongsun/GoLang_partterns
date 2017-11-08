package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Log("获取Builder")
	b := NewCar()
	t.Log("构建实例")
	b.Brand("BMW").Color("YELLOW").Name("MyCar").Show()
}
