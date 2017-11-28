# Go语言设计模式
主要的设计模式的Go语言实现。包含示例代码和单元测试。
Idea源自tmrts的开源项目[go-partterns](https://github.com/tmrts/go-patterns). tmrts的项目提供了设计模式中核心精炼代码。本项目将tmrts代码相关内容上翻译成中文并在此基础上结合开发经验进一步实现，更符合工程应用。

## 使用方法
下载代码
```
git clone https://github.com/DennisMao/gopartterns.git
```
测试
```
cd gopartterns/builder
go test -v
```


## 创建型模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|
| [抽象工厂](/creational/abstractFactory/abstractFactory.go) | 提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。 | ✔ |
| [建造者模式](/creational/builder/builder.go) | 通过使用多个简单的对象来逐步构建一个复杂的对象 | ✔ |
| [工厂模式](/creational/factory/factory.go) | 定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行。 | ✔ |
| [对象池模式](/creational/objectPool/objectPool.go) | 定义一个恒定大小对象组，该对小组能够给自行维护，用户可通过接口获取单个对象。 | ✔ |
| [单例模式](/creational/singleTon/singleTon.go) | 确保只有单个对象被创建，并提供了一种访问其唯一的对象的方式，可以直接访问，不需要实例化该类的对象。 | ✔ |
| [原型模式](/creational/prototype/prototype.go) | 用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。 | ✔ |
## 结构型模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|
| [桥接模式] | 将抽象部分与实现部分分离，使它们都可以独立的变化。 | ✘ |
| [适配器模式] | | ✘ |
| [装饰器模式](/structual/decorator/decorator.go) | 不改变原有结构情况下向现有对象增加新的功能 | ✔ |

## 行为型模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|

## 同步模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|

## 并发模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|
| [Wait模式](/concurrency/wait/wait.go) | 初始创建协程并注册waitgroup，在主线程等待所有子协程退出。 | ✔ |
| [Cancel模式](/concurrency/cancel/cancel.go) | 通过向协程中引入context变量，主线调用cancel向所有协程广播消息退出 | ✔ |


## 消息模式

| 模式 | 描述 | 完成状态 |
|:-------:|:----------- |:------:|
