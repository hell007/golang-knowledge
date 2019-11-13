package Bridge

import "fmt"

//创建桥接实现接口
type DrawApi interface {
	DrawCircle(radius int, x int, y int)
}

//创建实现了 DrawAPI 接口的实体桥接实现类
type RedCircle struct {
}

func (s *RedCircle) DrawCircle(radius int, x int, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

//创建实现了 DrawAPI 接口的实体桥接实现类
type GreenCircle struct {
}

func (s *GreenCircle) DrawCircle(radius int, x int, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

//使用 DrawAPI 接口创建抽象类 Shape。
type Shape struct {
	drawApi DrawApi
}

func (s *Shape) Shape(api DrawApi) {
	s.drawApi = api
}

//创建实现了 Shape 接口的实体类。
type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (s *Circle) Circle(x int, y int, radius int, drawAPI DrawApi) {
	s.shape.Shape(drawAPI)
	s.x = x
	s.y = y
	s.radius = radius
}

func (s *Circle) Draw() {
	s.shape.drawApi.DrawCircle(s.radius, s.x, s.y)
}