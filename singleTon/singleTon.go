/*
单例模式-懒汉式（未加锁）
核心解决：	一个全局使用的类，频繁创建和销毁
原理:		经常被全局使用的类，在内部自行创建实例，并且保证只有一次创建，
			对外提供访问这个唯一对象的接口，可以直接访问，无需再创建。
			在Go中，通过sync包中once功能保证创建实例的方法只会被调用一次
*/
package singleton

import (
	"errors"
	"sync"
)

type singleton struct {
	list map[string]string
}

///////////////  实例 //////////////
var (
	once     sync.Once
	instance singleton
)

///////////////  外部接口 //////////////
func GetInstance() singleton {
	once.Do(func() {
		//只初始化一次
		instance = singleton{
			list: make(map[string]string),
		}
	})
	return instance
}

//增加
func (s singleton) Add(key, value string) error {
	_, ok := s.list[key]
	if ok {
		return errors.New("value exists")
	}

	s.list[key] = value
	return nil
}

//删除
func (s singleton) Del(key string) {
	delete(s.list, key)
}

//获取
func (s singleton) Get(key string) (string, error) {
	v, ok := s.list[key]
	if !ok {
		return "", errors.New("key is not exists")
	}
	return v, nil
}
