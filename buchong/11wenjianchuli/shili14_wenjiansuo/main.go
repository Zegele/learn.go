// Go语言文件锁操作
// c.biancheng.net/view/5730.html
package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	_ "syscall"
	"time"
)

//使用Go语言开发一些程序的时候，往往出现多个进程同时操作同一份文件的情况
//这很容易导致文件中的数据混乱。
//这是我们就需要采用一些手段来平衡这些冲突，文件锁flock应运而生
//对于flock，最常见的例子就是Nginx,进程运行起来后就会把当前的PID写入这个文件
//当然如果这个文件已经存在了，也就是前一个进程还没有退出，
//那么Nginx就不会重新启动，所以flock还可以用来检测进程是否存在
//
//flock是对于整个文件的建议性锁
//如果一个进程在一个文件（inode）上放了锁
//其他进程是可以知道的（建议性锁不强求进程遵守）
//最棒的一点是，它的第一个参数是文件描述符，在此文件描述符关闭时，锁会自动释放
//而当进程终止时，所有的文件描述符均会被关闭，
//所以很多时候就不用考虑类似原子锁的事情
//
//示例：

// 文件锁
type FileLock struct {
	dir string
	f   *os.File
}

func New(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}

// 加锁
func (l *FileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f

	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

// 释放锁
func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}

func main() {
	test_file_path, _ := os.Getwd()
	locked_file := test_file_path

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			flock := New(locked_file)
			err := flock.Lock()
			if err != nil {
				wg.Done()
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("output: %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	time.Sleep(2 * time.Second)
}

// 在Windows系统下运行会出现下面错误：
// undefined: syscall.Flock
// undefined: syscall.LOCK_EX
// undefined: syscall.LOCK_NB
// undefined: syscall.Flock
// undefined: syscall.LOCK_UN

//是因为Windows不支持pid锁，所以需要在Linux或Mac下才能正常运行上面的程序。

// 上面代码中演示了同时启动10个goroutine，
//但在程序运行过程中，只有一个goroutine能获得文件锁
//其他的goroutine在获取不到flock后，会抛出异常的信息
//这样即可达到同一个文件在指定的周期内只允许一个进程访问的效果
//
//代码中文件锁的具体调用：
//syscall.Flock(int(f.Fd()),syscall.LOCK_EX|syscall.LOCK_NB)
//syscall.LOCK_EX, syscall.LOCK_NB 是什么意思？
//flock属于建议性锁，不具备强制性
//一个进程使用flock将文件锁住，另一个进程可以直接操作正在被锁的文件，修改文件中的数据
//原因在于flock只是用于检测文件是否被加锁，针对文件已经被加锁，另一个进程写入数据的情况
//内核不会阻止这个进程的写入操作，也就是建议性锁的内核处理策略
//
//flock主要三种操作类型
//LOCK_SH:共享锁，多个进程可以使用同一把锁，常被用作共享锁
//LOCK_EX:排他锁，同时只允许一个进程使用，常被用作写锁
//LOCK_UN:释放锁
//
//进程使用flock尝试锁文件时，如果文件已经被其他进程锁住，进程会被阻塞直到锁被释放，
//或调用flock的时候，采用LOCK_NB参数
//在尝试锁住该文件的时候，发现已经被其他服务锁住，返回错误，错误码：EWOULDBLOCK
//
//flock锁的释放非常具有特色，即可调用LOCK_UN参数来释放文件锁
//也可以通过关闭fd的方式来释放文件锁（flock的第一个参数时fd），
//意味着flock会随着进程的关闭而被自动释放掉。
