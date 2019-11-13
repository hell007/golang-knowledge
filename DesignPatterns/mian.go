package main

import (
	"./AdapterPattern"
	"./Bridge"
	"./BuilderPattern"
	"./FactoryPattern"
	"./FilterPattern"
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

//建造者
func testBuilderPattern() {
	builder := new(BuilderPattern.ComputerBuilder)
	director := BuilderPattern.Director{Builder: builder}
	computer := director.Create("I7", "32G", "4T")
	fmt.Println(*computer)
}

//适配器
func testAdapterPattern() {
	audioPlayer := AdapterPattern.AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}

//过滤器
func testFilterPattern() {
	var persons []FilterPattern.Person
	persons = append(persons, FilterPattern.GetPerson("Robert", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("John", "Male", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Laura", "Female", "Married"))
	persons = append(persons, FilterPattern.GetPerson("Diana", "Female", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Mike", "Male", "Single"))
	persons = append(persons, FilterPattern.GetPerson("Bobby", "Male", "Single"))

	male := new(FilterPattern.CriteriaMale)
	fmt.Println(male.MeetCriteria(persons))

	female := new(FilterPattern.CriteriaFemale)
	fmt.Println(female.MeetCriteria(persons))

	single := new(FilterPattern.CriteriaSingle)
	fmt.Println(single.MeetCriteria(persons))

	singleMale := new(FilterPattern.AndCriteria)
	singleMale.AndCriteria(single, male)
	fmt.Println(singleMale.MeetCriteria(persons))

	singleFemale := new(FilterPattern.AndCriteria)
	singleFemale.AndCriteria(single, female)
	fmt.Println(singleFemale.MeetCriteria(persons))
}

// 桥接
func testBridge() {
	redCircle := Bridge.Circle{}
	redCircle.Circle(100, 100, 10, &Bridge.RedCircle{})
	greenCircle := Bridge.Circle{}
	greenCircle.Circle(100, 100, 10, &Bridge.GreenCircle{})
	redCircle.Draw()
	greenCircle.Draw()
}

func main() {
	//测试
	testBridge()
}
