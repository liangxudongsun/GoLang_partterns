package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Log("实例化Builder")

	t.Log("构建实例")
	car := NewBuilder().Brand("BMW").Color("YELLOW").Name("MyCar").Build()
	t.Log("实例展示")
	car.Show()
}
