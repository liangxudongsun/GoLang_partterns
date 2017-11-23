package decorator

import (
	"testing"
)

func TestDecorate(t *testing.T) {

	oj := &TargetObject{
		Color: "yellow",
	}
	t.Log("源方法执行结果:" + oj.Draw())

	DrawDecorate(f)
	t.Log("目标修饰器执行结果:")
	DrawObjectDecorate(oj)

	k := DrawDecorate(Draw())
	t.Log("方法修饰器执行结果:", k)

}
