package main

import (
	"github.com/google/wire"
)

func main() {
	//运行
	InitializeAllInstance()
	//setup(context.Background())
	//beego.run()
}

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
	foo *Foo
}

func NewBar(foo *Foo) *Bar {
	return &Bar{
		foo: foo,
	}
}

type Instance struct {
	Foo *Foo
	Bar *Bar
}

var SuperSet = wire.NewSet(NewFoo, NewBar)

//go:generate wire
func InitializeAllInstance() *Instance {
	wire.Build(SuperSet, Instance{})
	return &Instance{}
}

/*
参考beego框架
----
 |---api 协议
 |---conf 配置
 |---cmd 启动脚本
 |---controllers
 |---dao   数据层
 |---router  路由
 |---service   业务层
 |---main.go
 |---go.mod
 |---README.md
*/
