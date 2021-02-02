package DecoratorPattern

import "fmt"

/**
意图：动态地给一个对象添加一些额外的职责。就增加功能来说，装饰器模式相比生成子类的方式更为灵活。

主要解决：一般的，我们为了扩展一个类经常使用继承的方式实现，由于继承为类引用静态特征，并且随着扩展功能的增加，子类会很庞大。

何时使用：在不向增加很多子类的情况下扩展类。

如何解决：将具体功能职责划分，同时继承装饰者模式。

关键代码：

Component 类充当抽象角色，不应该具体实现。
修饰类引用和继承 Component 类，具体扩展类重写父类方法。
优点：装饰结构和被装饰结构可以独立发展，不会相互耦合，装饰模式是继承的一个替代模式，装饰模式可以动态扩展一个实现类的功能。

缺点：多层装饰比较复杂

使用场景：

扩展一个类的功能
动态增加功能，动态撤销
 */

type Shape interface {
	Draw1()
}

type Rectangle struct {
}

func (r *Rectangle) Draw1() {
	fmt.Println("Shape: Rectangle")
}

type Circle struct {
}

func (c *Circle) Draw1() {
	fmt.Println("Shape: Circle")
}

type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) ShapeDecorator(decoratedShape Shape) {
	s.decoratedShape = decoratedShape
}

func (s *ShapeDecorator) Draw1() {
	s.decoratedShape.Draw1()
}

type RedShapeDecorator struct {
	shapeDecorator ShapeDecorator
}

func (s *RedShapeDecorator) RedShapeDecorator(decoratedShape Shape) {
	s.shapeDecorator.ShapeDecorator(decoratedShape)
}

func (s *RedShapeDecorator) Draw1() {
	s.shapeDecorator.Draw1()
	s.setRedBorder(s.shapeDecorator.decoratedShape)
}

func (s *RedShapeDecorator) setRedBorder(decoratedShape Shape) {
	fmt.Println("Border Color: Red")
}
