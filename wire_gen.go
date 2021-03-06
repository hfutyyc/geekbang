// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from week4.go:

//go:generate wire
func InitializeAllInstance() *Instance {
	foo := NewFoo()
	bar := NewBar(foo)
	instance := &Instance{
		Foo: foo,
		Bar: bar,
	}
	return instance
}

// week4.go:

func main() {

	InitializeAllInstance()

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
