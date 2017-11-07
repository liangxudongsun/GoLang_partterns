package singleton

import (
	"testing"
)

func TestSingleTon(t *testing.T) {
	//创建公共实例
	t.Log("创建单例")
	single := New()
	//添加数据
	t.Log("添加数据key=aa,value=123")
	single.Add("aa", "123")

	t.Log("查询key=aa")
	v, _ := single.Get("aa")
	t.Log("value=", v)
}
