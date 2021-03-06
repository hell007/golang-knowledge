## Golang-数据库知识


### 1、mysql驱动

[go-sql-driver](https://github.com/go-sql-drive/mysql)

[GoLand配置MySQL数据库](https://www.cnblogs.com/NolaLi/p/10738317.html)

#### go-xorm

[go-xorm](https://github.com/go-xorm/xorm)

[golang-xorm库快速学习](https://www.jianshu.com/p/13d46e5d2d2a)

需要下载依赖 go-xorm/ 

      git clone https://github.com/go-xorm/xorm.git
      git clone https://github.com/go-xorm/builder.git
      git clone https://github.com/go-xorm/core.git
      git clone https://github.com/go-xorm/cmd.git

- 关于 xorm

xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。我在项目中经常使用，它的特性如下、

    支持Struct和数据库表之间的灵活映射，并支持自动同步表结构
    事务支持
    支持原始SQL语句和ORM操作的混合执行
    使用连写来简化调用
    支持使用Id, In, Where, Limit, Join, Having, Table, Sql, Cols等函数和结构体等方式作为条件
    支持级联加载Struct
    支持LRU缓存(支持memory, memcache, leveldb, redis缓存Store) 和 Redis缓存
    支持反转，即根据数据库自动生成xorm的结构体
    支持事件
    支持created, updated, deleted和version记录版本（即乐观锁）
    
- xorm 工具

xorm 是一组数据库操作命令的工具，包含如下命令：

    reverse 反转一个数据库结构，生成代码
    shell 通用的数据库操作客户端，可对数据库结构和数据操作
    dump Dump数据库中所有结构和数据到标准输出
    source 从标注输入中执行SQL文件
    driver 列出所有支持的数据库驱动
    那我们该如何使用 reverse 命令根据数据表结构生成 go 代码呢？

go get github.com/go-xorm/cmd/xorm

go get github.com/go-xorm/xorm

到GOPATH\src\github.com\go-xorm\cmd\xorm 目录下，执行

    go build
    
这时在此目录了下生成xorm.exe文件

goxorm配置

github.com\go-xorm\cmd\xorm\templates\goxorm\config 进行配置
    
    lang=go
    genJson=1
    prefix=jie_ //重要


接下来在xorm/目录下，开始执行

    目录 /usr/local/go/extends/src/github.com/go-xorm/cmd/xorm

    ./xorm reverse mysql root:admin@/jie?charset=utf8 ./templates/goxorm

在/xorm/models/下生成

