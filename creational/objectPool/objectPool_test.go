package objectPool

import (
	"testing"
)

func TestObjectPool(t *testing.T) {
	t.Log("创建对象池")
	pool := NewPool(10)

	t.Log("获取对象")
	err, ob := pool.Get()
	if err != nil {
		t.Log(err)
	}

	ob.Brand = "AUDI"
	ob.Shape = "CIRCLE"
	t.Log("对象内容:", ob.Show())

	t.Log("关闭对象")
	ob.Close()

}
