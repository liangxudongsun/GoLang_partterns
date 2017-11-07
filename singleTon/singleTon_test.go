package singleton

import (
	"testing"
)

func TestSingleTon(t *testing.T) {
	t.Log("获取单例")
	single := GetInstance()
	//添加数据
	t.Log("添加数据key=aa,value=123")
	single.Add("aa", "123")

	t.Log("查询key=aa")
	v, _ := single.Get("aa")
	t.Log("value=", v)
}
