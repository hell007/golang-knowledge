package FactoryPattern

/*
意图：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。
主要解决：主要解决接口选择的问题。
何时使用：系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。
 */

import "fmt"

// struct
type AbsFactory struct {
}

// interface
type AbstractFactory interface {
	GetColor(c string) Color
}

type Color interface {
	Fill()
}

type Red struct {
}

func (s *Red) Fill() {
	fmt.Println("Red Fill")
}

type Green struct {
}

func (s *Green) Fill() {
	fmt.Println("Green Fill")

}

// func
func (f *AbsFactory) GetColor(c string) Color {
	if c == "" {
		return  nil
	}

	switch c {
	case "red":
		return new(Red)
	case "green":
		return new(Green)
	default:
		return nil
	}
}

