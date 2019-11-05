package main

import(
	"./FactoryPattern"
	"./SingletonPattern"
	"fmt"
)

// 抽象工厂
func testAbstractFactory() {
	f := new(FactoryPattern.AbsFactory)
	color := f.GetColor("red")
	color.Fill()
}

//单例
func testSingleton() {
	s1 := SingletonPattern.GetInstance1()
	s1.Count = 5
	fmt.Println(s1)
	s2 := SingletonPattern.GetInstance1()
	fmt.Println(s2)

	s3 := SingletonPattern.GetInstance5()
	s3.Count = 3
	fmt.Println(s3)
	s4 := SingletonPattern.GetInstance5()
	fmt.Println(s4)
}

func main() {
	//测试
	testSingleton()
}
