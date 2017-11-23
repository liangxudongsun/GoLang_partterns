/*
设计模式-Go语言
开发者:Razil
*/
package main

import (
	"log"

	"gopartterns/decorator"
)

func main() {
	log.Println("DesginParttern in Golang")

}

func k() {
	oj := &decorator.TargetObject{
		Color: "yellow",
	}
	log.Println("源方法执行结果:" + oj.Draw())

	log.Println("目标修饰器执行结果:")
	decorator.DrawObjectDecorate(oj)

	k := decorator.DrawDecorate(decorator.Draw())
	log.Println("方法修饰器执行结果:", k)
}
