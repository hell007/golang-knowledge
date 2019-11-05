package SingletonPattern

/*
注意：
1、单例类只能有一个实例。
2、单例类必须自己创建自己的唯一实例。
3、单例类必须给所有其他对象提供这一实例。

意图：保证一个类仅有一个实例，并提供一个访问它的全局访问点。
主要解决：一个全局使用的类频繁地创建与销毁。
何时使用：当您想控制实例数目，节省系统资源的时候。
如何解决：判断系统是否已经有这个单例，如果有则返回，如果没有则创建。
 */

import "sync"

type Single struct {
	Count int
}

var (
	single *Single
	once sync.Once
)

// --------------懒汉----------------------------------
func GetInstance1() *Single {
	if single == nil {
		single = new(Single)
	}
	return single
}
//懒汉模式存在线程安全问题，在第3步的时候，如果有多个线程同时调用了这个方法，
//那么都会检测到instance为nil,就会创建多个对象，所以出现了饿汉模式



// ----------------饿汉----------------------------------
func init() {
	single = new(Single)
}

func GetInstance2() *Single {
	return single
}
//饿汉模式将在包加载的时候就创建单例对象，当程序中用不到该对象时，浪费了一部分空间
//和懒汉模式相比，更安全，但是会减慢程序启动速度，所以我们可以在进一步修改程序

// 目前的读取yaml文件大多是这种模式



// ----------------------双锁-----------------------------
var lock *sync.Mutex = &sync.Mutex{}
func GetInstance3() *Single {
	lock.Lock()
	defer lock.Unlock()
	if single == nil {
		single = new(Single)
	}
	return single
}
//不过对于单例模式来说Golang在标准包中已经有了更好的解决方法: Sync.Once


// ---------------双锁进阶 还可使用原子load以及赋值----------
func GetInstance4() *Single {
	if single == nil {
		lock.Lock()
		defer lock.Unlock()
		single = new(Single)
	}
	return single
}

// --------------- sync.Once --------------------------------
func GetInstance5() *Single {
	once.Do(func() {
		single = new(Single)
	})
	return single
}