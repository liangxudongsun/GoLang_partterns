/*
设计模式-Go语言
开发者:Razil
主页: razil.cc
*/
package main

import (
	"log"

	ab "dev_designpattern/singleton"
)

func main() {
	log.Println("DesginParttern in Golang")

	t := ab.New()
	t.Add("a", "b")

	log.Println("")
}
