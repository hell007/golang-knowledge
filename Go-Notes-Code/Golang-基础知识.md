
## Golang-基础知识

### 常量、变量
	
	// 常量
	const i int = 10
	const a = "常量"

	// 定义单个变量
	// var v type = val

	var v1 string = "变量1"
	var v2 = "变量2"


	// 定义多个变量
	var v3, v4 string = "v3", "v4"
	var v3, v4 = "v3", "v4"


	// 简写
	v5, v6 := "v5", "v6"


### 字符串

a、字符串都是采用UTF-8字符集编码。字符串是用一对双引号（""）或反引号（` `）括起来定义，它的类型是string。

b、字符串是不可变的，但可进行切片操作;

c、声明一个多行的字符串可以通过` `来声明;

	
	func main() {
	
		// 修改
		s := "cello"
	
		fmt.Printf("%s\n", s)
	
		s = "h" + s[1:] // 字符串虽不能更改，但可进行切片操作
	
		fmt.Printf("%s\n", s)
	
		// +
		s1 := "hello,"
	
		s2 := " world"
	
		str := s1 + s2
	
		fmt.Printf("%s\n", str)
	
		// 如果要声明一个多行的字符串可以通过`来声明：
	
		m := `hello 
			world`
	
		fmt.Println(m)
	}


### 数值类型

a、整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称;

b、浮点数的类型有float32和float64两种（没有float类型），默认是float64;

c、go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。如果需要小一些的，也有complex64(32位实数+32位虚数)。复数的形式为RE + IMi，其中RE是实数部分，IM是虚数部分，而最后的i是虚数单位;

d、warn: 需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错;


	func main() {
		// 整数
		var a int8 = 1
		var b int8 = 3
		v1 := a + b
		fmt.Println(v1)
	
		var m int32 = 1
		var n int32 = 2
		v2 := m + n
		fmt.Println(v2)
	
		// 复数
		var v3 complex64 = 5 + 5i
		fmt.Printf("Value is: %v", v3)
	}


### 布尔类型

	var open bool      // 全局变量声明 一般声明
	var status = false // 忽略类型的声明

	func main() {
		success := true // 简短声明
		open = true // 赋值操作
	
		fmt.Println(open)
		fmt.Println(status)
		fmt.Println(success)
	}


### 派生类型

	(a) 指针类型（Pointer）
	(b) 数组类型
	(c) 结构化类型(struct)
	(d) Channel 类型
	(e) 函数类型
	(f) 切片类型
	(g) 接口类型（interface）
	(h) Map 类型


### 条件语句

注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

a、if
	
	func main() {
	
		x := 12
	
		if x > 10 {
			fmt.Println("x > 10")
		} else {
			fmt.Println("x < 10")
		}
	
		// Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	
		// 计算获取值y,然后根据y返回的大小，判断是否大于10。
		if y := 10; y > 10 {
			fmt.Println("y > 10")
		} else if y < 10 {
			fmt.Println("y < 10")
		} else {
			fmt.Println("y = 10")
		}
	
		//这个地方如果这样调用就编译出错了，因为y是条件里面的变量
		//fmt.Println(y)
	
	}

b、switch


	switch sExpr {
		case expr1:
			some instructions
		case expr2:
			some other instructions
		case expr3:
			some other instructions
		default:
			other code
	}

sExpr和expr1、expr2、expr3的类型必须一致;

Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true;

	func main() {
	
		// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		// 但是可以使用fallthrough强制执行后面的case代码
	
		a := 8
	
		switch a {
		case 7:
			fmt.Println("a=", 7)
		case 8:
			fmt.Println("a=", 8)
			fallthrough // 强制执行 case 9:
		case 9:
			fmt.Println("a=", 9)
		default:
			fmt.Println("a=", 0)
		}
	}



### 循环语句

	func main() {
	
		// 1、for
	
		a := 0
		for i := 0; i < 10; i++ {
			a += i
		}
		fmt.Println("a=", a)
	

		// 2、while
		//有些时候需要进行多个赋值操作，由于Go里面没有操作符，那么可以使用平行赋值i, j = i+1, j-1;
		//有些时候如果我们忽略expression1和expression3;
		//其中;也可以省略，那么就变成如下的代码了，这就是while的功能
	
		b := 1
		for b < 10 {
			b += b
		}
		fmt.Println("b=", b)
	

		// 3、在循环里面有两个关键操作break和continue,
		break操作是跳出当前循环，
		continue是跳过本次循环。
		// 当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置;
	
		// break和continue还可以跟着标号，用来跳到多重循环中的外层循环
	
		for j := 10; j > 0; j-- {
			if j == 6 {
				//break
				continue
			}
			fmt.Print(j)
	
			fmt.Println("")
		}
		// break打印出来10、9、8、7、6
		// continue打印出来10、9、8、7、6、4、3、2、1
	

		// 4、for配合range可以用于读取slice和map的数据
	
		fslice := make(map[string]string)
	
		fslice["one"] = "1"
	
		for k, v := range fslice {
			fmt.Println("fslice's key:", k) //注释会报错 处理如下；
			fmt.Println("fslice's val:", v)
		}
	
		// 由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,
		// 在这种情况下, 可以使用_来丢弃不需要的返回值 例如
	
		for _, v := range fslice {
			fmt.Println("fslice's val:", v)
		}
	}


### goto

 注意：请明智地使用它。用goto跳转到必须在当前函数内定义的标签会造成内存占用

	
	func main() {
	
		i := 0
	
	Here: //这行的第一个词，以冒号结束作为标签
	
		if i < 10 {
			println(i)
			i++
		}
	
		goto Here //跳转到Here去
	}


### iota枚举


Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1

	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)
	
	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
	
	const (
		h, i, j = iota, iota, iota //iota重置; h=0,i=0,j=0 iota在同一行值相同
	)
	
	const (
		a       = iota //重置; a=0
		b       = "B"
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g       = iota             //g = 4
	)
	
	func main() {
		fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	}


### 数组、多维数组


a、var arr [n]type
在[n]type中，n表示数组的长度，type表示存储元素的类型
 
b、长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。
 
c、数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型


	func main() {
	
		// 定义数值，赋值
		var arr [10]int // 声明了一个int类型的数组
	
		arr[0] = 42 // 数组下标是从0开始的
		arr[1] = 13 // 赋值操作

		fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
		fmt.Println(arr)
		fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0
	
		// :=
		a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	
		b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	
		c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	
		// 多维数组
		// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
		doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	
		// 上面的声明可以简化，直接忽略内部的类型
		easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	
		fmt.Println(doubleArray)
	
		fmt.Println(easyArray)
	}


### slice "动态数组"


	func main() {
	
		// 1、slice可以从一个数组或一个已经存在的slice中再次声明。slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	
		// 声明一个含有10个元素元素类型为int的数组
		var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
		// 声明两个含有int的slice
		var a, b []int
	
		// a指向数组的第3个元素开始，并到第五个元素结束，
		a = ar[2:5]
		//现在a含有的元素: ar[2]、ar[3]和ar[4]
	
		// b是数组ar的另一个slice
		b = ar[5:7]
		// b的元素是：ar[5]和ar[6]
	
		fmt.Println(a)
		fmt.Println(b)

	
		// 2、slice和array的对应关系图
	
		//slice的默认开始位置是0，ar[:n]等价于ar[0:n]
	
		//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
	
		//如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

	
		// 3、slice有几个有用的内置函数：
	
		//len 获取slice的长度
	
		fmt.Printf("长度为= %d \n" , len(a))
	
		//cap 获取slice的最大容量
	
		fmt.Printf("最大容量= %d \n", cap(a))
	
		//append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	
		c := append(a, 6)
		fmt.Println(c)
	
		//copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
		// 复制的元素依次替换 dist[0] dist[1] ... ,并且dist src 的len不变
	
		dist := a
		src := b
	
		n := copy(dist, src)
		fmt.Printf("复制了 %d 个元素 \n", n) //复制了多少个元素
		fmt.Println(dist)
		fmt.Println(src)
	
		// 创建slice
	
		sliceA := make(map[int]string)
		sliceA[0] = "曹操"
		sliceA[1] = "刘备"
		sliceA[2] = "孙权"
		fmt.Println(sliceA)
	
		for _,val := range sliceA {
			fmt.Println(val)
		}
	}


### range

Go语言中range关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。

在数组和切片中它返回元素的索引和索引对应的值，在集合中返回key-value对。


### map(集合)

map[keyType]valueType


	func main() {
	
		// 1、声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
		var numbers map[string]int
	
		// 另一种map的声明方式
		numbers = make(map[string]int)
		numbers["one"] = 1  //赋值
		numbers["ten"] = 10 //赋值
		numbers["three"] = 3
	
		fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
		// 打印出来如:第三个数字是: 3
	
		//map过程中需要注意的几点：
		//a、map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取;
		//b、map的长度是不固定的，也就是和slice一样，也是一种引用类型;
		//c、内置的len函数同样适用于map，返回map拥有的key的数量；
		//d、map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11;
		//e、map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制;
	

		// 2、map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	
		// 初始化一个字典
		rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
		val, success := rating["Go"]
		if success {
			fmt.Println("Go=", val)
		} else {
			fmt.Println("Go=", "not found")
		}
	

		//3、通过delete删除map的元素
	
		delete(rating, "C") // 删除key为C的元素
		fmt.Println(rating)

	
		// 4、map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	
		m := make(map[string]string)
		m["key"] = "val"

		n := m
		n["key"] = "value"
	
		fmt.Println("m=", m)
		fmt.Println("n=", n)
	

		// 5、make、new操作
		//make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
	
		// new返回指针
		// make返回初始化后的（非零）值。
	

		// 6、零值
		// 关于“零值”，所指并非是空值，而是一种“变量未填充前”的默认值，通常为0。 此处罗列 部分类型 的 “零值”
	
		/*
			int     0
			int8    0
			int32   0
			int64   0
			uint    0x0
			rune    0 //rune的实际类型是 int32
			byte    0x0 // byte的实际类型是 uint8
			float32 0 //长度为 4 byte
			float64 0 //长度为 8 byte
			bool    false
			string  ""
		*/
	}


### struct类型


a、struct定义

	// 声明一个新的类型
	type person struct {
		name string
		age  int
	}

除了上面这种P的声明使用之外，还有另外几种声明使用方式：

1）、按照顺序提供初始化值

	P := person{"Tom", 25}


2）、通过field:value的方式初始化，这样可以任意顺序

	P := person{age:24, name:"Tom"}


3）、当然也可以通过new函数分配一个指针，此处P的类型为*person

	P := new(person)


b、struct的匿名字段

struct定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

	type Human struct {
		name   string
		age    int
		weight int
	}
	
	type Student struct {
		Human // 匿名字段，那么默认Student就包含了Human的所有字段
		sex   string
	}

	func main() {
	
		// 1.
		var p person
		p.name = "曹操"
		p.age = 23
		fmt.Printf("我叫%s，今年%d岁了，是该君临天下了吧！", p.name, p.age)
	
		// 2.
	
		// 定义一个学生
		joy := Student{Human{"joy", 22, 110}, "男"}
		fmt.Println("姓名：", joy.name)
		fmt.Println("年龄：", joy.age)
		fmt.Println("体重：", joy.weight)
		fmt.Println("性别：", joy.sex)
	
		// 修改信息
		joy.weight = 120
		fmt.Println("变胖了，现在都：", joy.weight)
	
		joy.sex = "保密"
		fmt.Println("注重隐私，性别：", joy.sex)
	}



### reflect

Go语言实现了反射，所谓反射就是能检查程序在运行时的状态

 	 t := reflect.TypeOf(i)

得到类型的元数据,通过t我们能获取类型定义里面的所有元素

	 v := reflect.ValueOf(i) 

得到实际的值，通过v我们获取存储在里面的值，还可以去改变值


	import (
		"fmt"
		"reflect"
		"time"
	)
	
	type Admin struct {
		Id     int
		Name   string
		Phone  string
		Status int
		Time   int64
	}
	
	func main() {
	
		//初始化
		a := Admin{Id: 0, Name: "小曹", Phone: "13888888887", Status: 1, Time: 0}
		fmt.Println("====", a)
	
		//get
		v := reflect.ValueOf(a)
		name := v.FieldByName("Name").String()
		fmt.Println("===", name)
	
		//set
		b := reflect.ValueOf(&a).Elem()
		b.FieldByName("Id").SetInt(1000)
		b.FieldByName("Name").SetString("曹操")
		b.FieldByName("Phone").SetString("138888888888")
		b.FieldByName("Time").SetInt(time.Now().Unix())
		fmt.Println("====", a)
	
		//save
		admin := Admin{}
		admin.Id = 2000
	
		fmt.Println("admin===", admin)
	}



### func 函数

	func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
	}

1）、关键字func用来声明一个函数funcName;

2）、函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔;

3）、函数可以返回多个值;

4）、上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型;

5）、如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号;

6）、如果没有返回值，那么就直接省略最后的返回信息;

7）、如果有返回值， 那么必须在函数的外层添加return语句;


a、定义函数

	func max(a, b int) int {
		if a > b {
			return a
		}
		return b
	}


b、多个返回值

	func count(a, b int) (sum int, multiplied int) {
		sum = a + b
		multiplied = a * b
		return sum, multiplied
	}


c、Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参。

arg ...int告诉Go这个函数接受不定数量的参数。注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice：

	func myfunc(arg ...int) {
	
		for _, val := range arg {
			fmt.Printf("arg the number is: %d\n", val)
		}
	}


d、传值与传指针

当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，
当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上.

	// 传值
	func test(a int) int {
		a = a + 1 //更改 a
		return a
	}

	// 传指针
	func test2(a *int) int {
		*a = *a + 1 //更改 a
		return *a
	}

传指针使得多个函数能操作同一个对象。

传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。

如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。

Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。

（注：若函数需改变slice的长度，则仍需要取地址传递指针）

e、函数作为值、类型

Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是拥有相同的参数，相同的返回值的一种类型


	type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])


	type funcInt func(int) bool // 声明了一个函数类型
	
	func odd(integer int) bool {
		if integer%2 == 0 {
			return false
		}
		return true
	}
	
	func even(integer int) bool {
		if integer%2 == 0 {
			return true
		}
		return false
	}
	
	// 声明的函数类型在这个地方当做了一个参数
	func filter(slice []int, f funcInt) []int {
		var result []int
	
		for _, val := range slice {
			if f(val) {
				result = append(result, val)
			}
		}
		return result
	}

总结：函数当做值和类型在我们写一些通用接口的时候非常有用，通过上面例子我们看到funcInt这个类型是一个函数类型，然后两个filter函数的参数和返回值与funcInt类型是一样的，但是我们可以实现很多种的逻辑，这样使得我们的程序变得非常的灵活。


f、Panic和Recover 内建函数

	var user = os.Getenv("USER")

	func init() {
		if user == "" {
			panic("no value for $USER")
		}
	}

	func throwsPanic(f func()) (b bool) {
		defer func() {
			if x := recover(); x != nil {
				b = true
			}
		}()
		f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
		return
	}

	func main() {
	
		// 1.
		max_xy := max(1, 2)
		fmt.Println("max val=", max_xy)
	
		// 2.
		s, m := count(3, 4)
		fmt.Println("3 + 4 =", s)
		fmt.Println("3 * 4 =", m)
	
		// 3.
		myfunc(1, 2, 3, 4, 5)
	
		// 4.
		x := 4
		fmt.Println("x = ", x) // 应该输出 "x = 4"
		x1 := test(x) //调用test(x)
		fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 5"
		fmt.Println("x = ", x) // 应该输出"x = 4"
	
		y := 4
		y1 := test2(&y) //调用test2(y)
		fmt.Println("y+1 = ", y1) // 应该输出"y+1 = 5"
		fmt.Println("y = ", y) // 应该输出"y = 5"
	
		// 5.
		slice := []int{1, 2, 3, 4, 5, 6}
		fmt.Println("slice = ", slice)
		// 函数当做值来传递了
		sliceOdd := filter(slice, odd)
		fmt.Println("odd val = ", sliceOdd)
	
		// 函数当做值来传递了
		sliceEven := filter(slice, even)
		fmt.Println("even val = ", sliceEven)
	
		// 6.
		//init()
		//throwsPanic()
	}


### method 面向对象

带有接收者的函数，我们称为method
 
	func (r ReceiverType) funcName(parameters) (results)

在使用method的时候重要注意几点:

1）、虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样；

2）、method里面可以访问接收者的字段；

3）、调用method 通过.访问，就像struct里面访问字段一样；



a、method使用

	// 接收者
	type Rectangle struct {
		width, height float64
	}
	
	type Circle struct {
		radius float64
	}
	
	// method
	func (r Rectangle) area() float64 {
		return r.width * r.height
	}
	
	func (c Circle) area() float64 {
		return math.Ceil(c.radius * c.radius * math.Pi)
	}


b、method继承

	// 接收者
	type Human struct {
		name string
		age  int
		sex  string
	}
	
	type Student struct {
		Human  // 匿名字段
		school string
	}
	
	type Employee struct {
		Human //匿名字段
		job   string
	}
	
	// Human 定义method
	func (h *Human) Introduce() {
		fmt.Printf("你好，我叫%s，今年%d岁。\n", h.name, h.age)
	}


c、method重写

	// Employee的method重写Human的method
	func (e *Employee) Introduce() {
		fmt.Printf("你好，我叫%s，今年%d岁，从事%s工作！\n", e.name, e.age, e.job)
	}

Go里面的面向对象是如此的简单，没有任何的私有、公有关键字，通过大小写来实现(大写开头的为公有，小写开头的为私有)，方法也同样适用这个原则。
	
	func main() {
		// 1.
		r := Rectangle{12, 7.8}
		fmt.Println("长方形的面积 = ", r.area())
	
		c := Circle{3}
		fmt.Println("圆的面积 = ", c.area())
	
		// 2.
		jay := Student{Human{"jay", 25, "男"}, "北大"}
		jay.Introduce()
	
		// 3.
		tom := Employee{Human{"tom", 26, "女"}, "工程师"}
		tom.Introduce()
	}



### interface

interface是一组method签名的组合，我们通过interface来定义对象的一组行为


a、interface类型：

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

	// 接收者
	type Human struct {
		name string
		age  int
		sex  string
	}
	
	type Student struct {
		Human  //匿名字段Human
		school string
	}
	
	type Employee struct {
		Human  //匿名字段Human
		job    string
		salary float32
	}
	
	// method
	func (h Human) Introduce() {
		fmt.Printf("你好，我叫%s，今年%d岁。\n", h.name, h.age)
	}
	
	func (h Human) Like(lyrics string) {
		fmt.Println("我喜欢唱歌", lyrics)
		fmt.Printf("\n")
	}
	
	func (e Employee) Introduce() {
		fmt.Printf("你好，我叫%s，今年%d岁，从事%s工作！\n", e.name, e.age, e.job)
	}

定义interface

	type Mine interface {
		Introduce()
		Like(lyrics string)
	}

通过上面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现


b、空interface(interface{})

空interface不包含任何的method，正因为如此，所有的类型都实现了空interface。

空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。

	// 定义a为空接口
	var a interface{}


c、interface变量存储的类型

// Comma-ok断言

// switch测试


d、嵌入interface

	func main() {
		// 1.
		jay := Student{Human{"jay", 23, "男"}, "北大"}
		tom := Employee{Human{"Tom", 37, "男"}, "工程师", 5000}
		lily := Student{Human{"lily", 43, "女"}, "北大"}
	
		//定义Mine类型的变量i
		var i Mine

		//i能存储Student
		i = jay
		i.Introduce()
		i.Like("爱你一万年，期待你的表演！")
	
		//i也能存储Employee
		i = tom
		i.Introduce()
		i.Like("我是一只小小鸟！")
	
		//定义了slice Mine
		x := make([]Mine, 3)
		//这三个都是不同类型的元素，但是他们实现了interface同一个接口
		x[0], x[1], x[2] = jay, tom, lily
	
		for _, val := range x {
			val.Introduce()
		}
	
		// 2.
		fmt.Println("\na可以存储任意类型的数值")
		var b int = 5
		s := "Hello world"
		a = b
		fmt.Println("a的类型是int， a = ", a)
		a = s
		fmt.Println("a的类型是string，a = ", a)
	}


### error类型


Go内置有一个error类型，专门用来处理错误信息，Go的package里面还专门有一个包errors来处理错误。

	import (
		"errors"
		"fmt"
	)
	
	func main() {
		err := errors.New("出错了！")
		if err != nil {
			fmt.Println(err)
		}
	}



### 并发

a、goroutine

Go语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。goroutine（协程）是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：

	go 函数名( 参数列表 )

Go允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间


b、通道（channel）

通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

	ch <- v    // 把 v 发送到通道 ch
	v := <-ch  // 从 ch 接收数据
	           // 并把值赋给 v

注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。


- 用途：管道实现同步

这种方式是一种比较完美的解决方案， goroutine / channel 它们也是在 go 里面经常搭配在一起的一对.

	func main(){
		ch := make(chan struct{})
		count := 2 // count 表示活动的协程个数
	
		go func() {
			fmt.Println("Goroutine 1")
			ch <- struct{}{} // 协程结束，发出信号
		}()
	
		go func() {
			fmt.Println("Goroutine 2")
			ch <- struct{}{} // 协程结束，发出信号
		}()
	
		for range ch {
			// 每次从ch中接收数据，表明一个活动的协程结束
			count--
			// 当所有活动的协程都结束时，关闭管道
			if count == 0 {
				close(ch)
			}
		}
	}


- go 里面也提供了更简单的方式 —— 使用 sync.WaitGroup。

WaitGroup 顾名思义，就是用来等待一组操作完成的。WaitGroup 内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：

Add() 用来添加计数；

Done() 用来在操作结束时调用，使计数减一；

Wait() 用来等待所有的操作结束，即计数变为 0，该函数会在计数不为 0 时等待，在计数为 0 时立即返回；

	import (
		"fmt"
		"sync"
	)

	func main(){
		var wg sync.WaitGroup
		wg.Add(2) // 因为有两个动作，所以增加2个计数
	
		go func() {
			fmt.Println("Goroutine 1")
			wg.Done() // 操作完成，减少一个计数
		}()
	
		go func() {
			fmt.Println("Goroutine 2")
			wg.Done() // 操作完成，减少一个计数
		}()
	
		wg.Wait() // 等待，直到计数为0
	}


c、通道缓冲区（Buffered Channels）

通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：

	ch := make(chan int, value)

当 value = 0 时，channel 是无缓冲阻塞读写的，当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入

	func main() {
		// 这里我们定义了一个可以存储整数类型的带缓冲通道
		// 缓冲区大小为2
		ch := make(chan int, 2)
		
		// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
		// 而不用立刻需要去同步读取数据
		ch <- 1
		ch <- 2
		
		// 获取这两个数据
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}

- 总结：

channel使用的注意事项

	channel中只能存放指定的数据类型
	channle的数据放满后，就不能再放入了
	如果从channel取出数据后，可以继续放入
	在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock

- 注意：

如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。


d、Go 遍历通道与关闭通道（Range和Close）

上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，所以也可以通过range关键字来实现遍历读取到的数据，像操作slice或者map一样操作缓存类型的channel，类似于与数组或切片。


	func fibonacci(n int, c chan int) {
	    x, y := 0, 1
	    for i := 0; i < n; i++ {
	            c <- x
	            x, y = y, x+y
	    }
	    close(c)
	}
	
	func main() {
	    c := make(chan int, 10)
	    go fibonacci(cap(c), c)
	    // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	    // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	    // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	    // 会结束，从而在接收第 11 个数据的时候就阻塞了。
	    for i := range c {
	    	fmt.Println(i)
	    }
	}

for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

- 注意：

记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic;

另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的;


e、Select

上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。


select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。

	func fibonacci(c, quit chan int) {
	    x, y := 1, 1
	    for {
	        select {
	        case c <- x:
	            x, y = y, x + y
	        case <-quit:
	            fmt.Println("quit")
	            return
	        }
	    }
	}
	
	func main() {
	    c := make(chan int)
	    quit := make(chan int)
	    go func() {
	        for i := 0; i < 10; i++ {
	            fmt.Println(<-c)
	        }
	        quit <- 0
	    }()
	    fibonacci(c, quit)
	}

在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。

	select {
	case i := <-c:
	    // use i
	default:
	    // 当c阻塞的时候执行这里
	}


- 超时

有时候会出现goroutine阻塞的情况，那么如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：

	func main() {
	    c := make(chan int)
	    o := make(chan bool)
	    go func() {
	        for {
	            select {
	                case v := <- c:
	                    println(v)
	                case <- time.After(5 * time.Second):
	                    println("timeout")
	                    o <- true
	                    break
	            }
	        }
	    }()
	    <- o
	}


### defer 延迟语句

Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。

特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。

	import (
		"fmt"
		"os"
	)
	
	func main() {
	
		filePath := "D:/Dev/cygwin/work/golang-knowledge/Go-Notes-Code/Basic/city.txt"
		//1、打开文件
		//file, err := os.Open(filePath) // For read access.
	
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066)
	
		//2、关闭文件
		defer file.Close()
	
		if err != nil {
			fmt.Println("打开失败！")
		}
	
		//3、读取文件
		var b []byte = make([]byte, 4096)
	
		n, err := file.Read(b)
	
		if err != nil {
			fmt.Println("Open file Failed", err)
		}
	
		data := string(b[:n])
		fmt.Println(data)
	
	}


