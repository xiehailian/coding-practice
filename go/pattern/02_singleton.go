package pattern

import "sync"

/*
 * 饿汉式单例模式，类加载时就创建好了唯一实例，使用时性能高且线程安全
 * 缺点：因为类加载时就创建好了，如果类成员占用资源比较多，浪费内存空间
 */
type hungerSingleton struct{}

var hungerSingletonInstance = &hungerSingleton{}

func GetHungerSingletonInstance() *hungerSingleton {
	return hungerSingletonInstance
}

/*
 * 懒汉式单例模式，使用时再创建
 * 缺点：线程不安全，多个线程同时访问时可能创建多个单例
 */
type lazySingleton struct{}

var lazySingletonInstance *lazySingleton

func GetLazySingletonInstance() *lazySingleton {
	if lazySingletonInstance == nil {
		lazySingletonInstance = &lazySingleton{}
	}
	return lazySingletonInstance
}

/*
 * 懒汉加锁，加锁解决并发时的线程不安全问题
 * 缺点: 每次访问时都要加锁解锁，浪费性能
 */

type lazySingletonWithLock struct{}

var lazySingletonWithLockInstance *lazySingletonWithLock
var mu sync.Mutex

func GetLazySingletonWithLockInstance() *lazySingletonWithLock {
	mu.Lock()
	defer mu.Unlock()

	if lazySingletonWithLockInstance == nil {
		lazySingletonWithLockInstance = &lazySingletonWithLock{}
	}
	return lazySingletonWithLockInstance
}

/*
 * 懒汉双重锁，避免每次加锁，提高效率
 */

type lazySingletonWithDoubleLock struct{}

var lazySingletonWithDoubleLockInstance *lazySingletonWithDoubleLock
var lock sync.Mutex

func GetLazySingletonWithDoubleLockInstance() *lazySingletonWithDoubleLock {
	if lazySingletonWithDoubleLockInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if lazySingletonWithDoubleLockInstance == nil {
			lazySingletonWithDoubleLockInstance = &lazySingletonWithDoubleLock{}
		}
	}
	return lazySingletonWithDoubleLockInstance
}

/*
 * 推荐使用sync.Once
 */

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
