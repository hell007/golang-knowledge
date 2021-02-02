package FacadePattern

import "fmt"

/**
意图：为子系统中的一组接口提供一个一致的界面或调用方式，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。

主要解决：降低访问复杂系统时的复杂度，简化客户端与之的接口。

何时使用：

客户端不需要知道系统内部的复杂联系，整个系统只需要提供一个接待员即可。
可定义系统的入口
关键代码：在客户端和复杂系统之间加一层，这一层将调用顺序、依赖关系等处理好。

优点：

减少系统互相依赖
提高灵活性
提高了安全性
缺点：不符合开闭原则，如果要改东西很麻烦，继承重写都不合适。

使用场景：

为复杂的模块或子系统提供外界访问的模块
子系统相对独立
预防低水平人员带来的风险
 */

type Shape interface {
	Draw2()
}

type Rectangle struct {
}

func (r *Rectangle) Draw2() {
	fmt.Println("Rectangle::draw()")
}

type Square struct {
}

func (s *Square) Draw2() {
	fmt.Println("Square::draw()")
}

type Circle struct {
}

func (c *Circle) Draw2() {
	fmt.Println("Circle::draw()")
}
