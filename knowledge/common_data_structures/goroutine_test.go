//  goroutine_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/05.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

package commondatastructures_test

// 知识补充
/**
线程上下文切换：3.5us
协程上下文切换：120ns
线程大约是协程的 30倍
资料来源：https://zhuanlan.zhihu.com/p/80037638

goroutine 解释：
routine: 使计算机执行某一特定任务的指令列表；
简单概括一下就是，Go语言调度函数及其执行上下文的数据结构；

栈大小对比：
	- goroutine 的初始栈大小为 2k，根据用户使用内存情况自动增长或缩小；
	- Linux上的线程栈大小为 4k~10M，用户可自行修改；
	- iOS的线程栈是固定的，APP主线程 1M，用户创建的线程 512K，不同系统版本可能会有变化；

*/

/**
runtime.NumCPU()       // 返回当前CPU内核数
runtime.GOMAXPROCS(2)  // 设置运行时最大可执行CPU数
runtime.NumGoroutine() // 当前正在运行的goroutine 数
*/
