/*
意图：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。
主要解决：在运行期建立和删除原型。

何时使用： 
1、当一个系统应该独立于它的产品创建，构成和表示时。
2、当要实例化的类是在运行时刻指定时，例如，通过动态装载。
3、为了避免创建一个与产品类层次平行的工厂类层次时。
4、当一个类的实例只能有几个不同状态组合中的一种时。

建立相应数目的原型并克隆它们可能比每次用合适的状态手工实例化该类更方便一些。

如何解决：利用已有的一个原型对象，快速地生成和原型对象一样的实例。

优点： 1、性能提高。 2、逃避构造函数的约束。
缺点： 1、配备克隆方法需要对类的功能进行通盘考虑，这对于全新的类不是很难，但对于已有的类不一定很容易，
特别当一个类引用不支持串行化的间接对象，或者引用含有循环结构的时候
*/

package PrototypePattern

import (
	"bytes"
	"encoding/gob"
)

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

//内存数量以及大小
type Memory struct {
	Count      int
	MemorySize []int
}

//电脑信息
type Computer struct {
	SystemName string              //系统名字
	UseNumber  int                 //使用次数
	Memory     Memory              //存储
	Fan        map[string]FanSpeed //风扇
	Money      Money               //售价
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
