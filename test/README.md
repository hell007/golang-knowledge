
## go modules

关于go的第三方包的依赖管理，项目小的时候还不觉得多麻烦，当项目越来越大，引入的第三方包越来越多的时候，
这个时候问题就很头痛了，golang官方建议是一个项目一个GOPATH，但是鉴于很多人都不这么做（我也不这么做，因为麻烦），
于是在go1.11的时候推出了这个go modules来解决依赖管理的问题

- 使用

由于这个go modules和GOPATH可能是天生不和，如果配置了GOPATH的话，这个go modules是默认不开启的。
开启也很简单，只要在命令行中键入下面命令就行了。

    export GO111MODULE=on	# 开启GoModule特性

不过这种是一次性的，重启之后就会消失，所以可以在环境信息文件/etc/profile的最后面添加上一行

    export GO111MODULE=on

使用命令source /etc/profile让其生效。


GO111MODULE解释, 当为on时则使用Go Modules,go 会忽略 $GOPATH和 vendor文件夹,只根据go.mod下载依赖。
当为 off时则不适用新特性 Go Modules支持，它会查找 vendor目录和 $GOPATH来查找依赖关系，也就是继续使用“GOPATH模式”。
当为 auto时或未设置时则根据当前项目目录下是否存在 go.mod文件或 $GOPATH/src之外并且其本身包含go.mod文件时才会使用新特性 
Go Modules模式，并且auto为 GO111MODULE的默认值。

## goproxy

关于goproxy，简单来说就是一个代理，让我们更方便的下载哪些由于墙的原因而导致无法下载的第三方包，
比如golang.org/x/下的包，虽然也有各种方法解决，但是，如果是你在拉取第三方包的时候，而这个包又依赖于golang.org/x/下的包，
你本地又恰恰没有，当然不嫌麻烦的话，也可以先拉取golang.org/x/下的包，再拉取第三方包。

这个goproxy强大地方就在于代理，而它官网是这样介绍自己的【A Global Proxy for Go Modules】，
就是这么强大，全球代理，让世界没有难下的包


- 使用

一般来说，goproxy是和go modules配合使用的。
首先要确保go modules是开启的，如果没有开启可以在命令行键入下面命令即可

    export GO111MODULE=on
    export GOPROXY=https://goproxy.io

不过这种是一次性的，重启之后就会消失，所以可以在环境信息文件/etc/profile的最后面添加上下面这两行

    export GO111MODULE=on
    export GOPROXY=https://goproxy.io

最后使用source /etc/profile让其生效。

## go.mod 和 go.sum介绍

go.mod是Go项目的依赖描述文件，该文件主要用来描述两个事情：

1、当前项目名(module)是什么。每个项目都应该设置一个名称，当前项目中的包(package)可以使用该名称进行相互调用。

2、当前项目依赖的第三方包名称。项目运行时会自动分析项目中的代码依赖，生成go.sum依赖分析结果，
随后go编译器会去下载这些第三方包，然后再编译运行。

go.sum依赖分析文件，记录每个依赖库的版本和哈希值

一般情况下，go.sum应当被添加到版本管理中随着go.mod文件一起提交

## go modules常用命令

    go mod init: 初始化modules
    go mod download: 下载依赖的module到本地cache
    go mod edit: 编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
    go mod graph: 以文本模式打印模块需求图
    go mod tidy: 检查，删除错误或者不使用的modules，以及添加缺失的模块
    go mod vendor: 生成vendor目录，将依赖复制到vendor目录下面
    go mod verify: 验证依赖是否正确
    go mod why： 解释为什么需要依赖

    go list -m： 查看主模块的路径
    go list -m -f={{.Dir}}： 查看主模块的根目录
    go list -m all： 查看当前的依赖和版本信息
    
    
## goland ide下使用go modules    

在goland下，是推荐使用goland配置vgo来快速使用go modules的。
而vgo是基于Go Modules规范的包管理工具，同官方的go mod命令工具类似。

- 下面是配置流程：

在设置页面里，根据下图来配置就可以了

![](./images/goland-module-1.png)

还有一种是新建项目时就想使用go modules，那么可以根据下面的图来配置，
而这种的话，新建时就会在项目根目录下自动创建一个 go.mod 文件

![](./images/goland-module-2.png)

接下来就以新建项目的这种情况来介绍怎么在goland下使用go modules

先创建一个项目，在创建页面的时候就选择"Go Modules(vgo)“这个选项，
强烈建议在"Proxy"这一栏填上"https://goproxy.io”，毕竟墙的存在

在创建完项目的时候，会在项目目录下自动生成一个go.mod文件，而这个文件初始的时候，内容就只有以下的两句话

    module testModuleProject		// 模块名，默认使用项目名。

    go 1.12		// 当前go的版本	

而模块名就是我们说的项目依赖路径，通常来说该地址就是项目的仓库地址，所有需要引用项目包的地址都要填写这个路径，
无论是内部之间引用还是外部引用。也就是说，在项目启用了 go modules 之后，引用包必须跟 go mod 文件第一行包名一样

好，接下来我们创建main.go，随便import一个包，会发现这个包是红色的，就算本地的$GOPATH/src/github.com/有这个第三方包
也一样会报红。这是因为我们启动了GO111MODULE =on，那么只会使用go module的方式寻找依赖包，
由于这时候go.mod里没有任何加入依赖的语句，那么自然是找不到，所以会报红。
最简单的方法就是把鼠标放在报红的地方，会弹出来解决方案，点击Sync package of testModuleProject或用快捷键"Alt+Shift+Enter"，
把包同步到模块里，其实这里的同步包就是相当于命令go mod download，只不过比较方便而已。

![](./images/goland-module-3.png)

![](./images/goland-module-4.png)

而且在go.mod里自动添加了一条语句，而这条语句就是重中之重。其中v1.2.0表示vgo检测到的最新版本。

go.mod文件可以通过require，replace和exclude语句的使用精确软件包集。

require语句指定的依赖项模块
replace语句可以替换依赖项模块
exclude语句可以忽略依赖项模块

![](./images/goland-module-5.png)

而依赖的包会被下载到$GOPATH/pkg/mod/下，我们可以在箭头所指位置能看到当前项目的所有依赖

![](./images/goland-module-6.png)

[文章来源](https://blog.csdn.net/qq_42403866/article/details/93654421)

## 问题

- 题外话

在https://goframe.org/prepare/vendor里看到的关于go modules的版本选择算法的简单介绍

版本选择算法
当项目中存在同一个第三方包依赖，并且依赖版本不一致时，Go Modules使用的“最小版本选择算法”(The minimal version selection algorithm: https://github.com/golang/go/wiki/Modules#version-selection )。

例如，如果您的模块依赖于具有require D v1.0.0的模块A，并且您的模块还依赖于具有require D v1.1.1的模块B，则最小版本选择将会选择D的v1.1.1版本用以构建（使用最高版本）。

十分有意思的是，这个算法名字叫“最小版本选择算法”，然而内容却是“最高版本选择算法”，若有纠结于此的同学欢迎向官方提issue：https://github.com/golang/go/issues


##### 1、golang 依赖包被墙解决方法

1、使用 gopm 进行包管理

2、在 github 上找到对应的包地址，再使用 go get 安装

3、对于 1.11 以上版本，可以在 go.mod 中通过 replace 替换地址，下方是能满足iris、gin和gorm的依赖

    module gin

    go 1.12

    replace golang.org/x/net => github.com/golang/net latest
    replace golang.org/x/tools => github.com/golang/tools latest
    replace golang.org/x/crypto => github.com/golang/crypto latest
    replace golang.org/x/sys => github.com/golang/sys latest
    replace golang.org/x/text => github.com/golang/text latest
    replace golang.org/x/sync => github.com/golang/sync latest
    replace golang.org/x/exp => github.com/golang/exp latest
    replace golang.org/x/lint => github.com/golang/lint latest
    replace golang.org/x/oauth2 => github.com/golang/oauth2 latest
    replace golang.org/x/time => github.com/golang/time latest
    replace golang.org/x/image => github.com/golang/image latest
    replace golang.org/x/mobile => github.com/golang/mobile latest
    replace cloud.google.com/go => github.com/googleapis/google-cloud-go latest
    replace google.golang.org/grpc => github.com/grpc/grpc-go latest
    replace google.golang.org/appengine => github.com/golang/appengine latest
    replace google.golang.org/genproto => github.com/googleapis/go-genproto latest
    replace google.golang.org/api => github.com/googleapis/google-api-go-client latest
    
4、设置环境变量：GOPROXY=https://goproxy.io 可以使用代理，不仅能解决强，速度还很快！