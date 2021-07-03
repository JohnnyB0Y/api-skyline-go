# 管道的数据结构

```go
// 管道的数据结构
type custom_chan struct {
	qcount uint // 当前队列剩余的元素个数
	dataqsize uint // 队列的容量
	buf unsafe.Pointer // 环形队列指针
	elemsize uint16 // 元素的大小
	closed uint32 // 管道关闭状态
	elemtype *_type // 元素类型
	sendx uint // 下一个待写入元素的位置
	recvx uint // 下一个待读取元素的位置
	sendq waitq // 等待写入的协程队列
	recvq waitq // 等待读取的协程队列
	lock mutex // 互斥锁
}
```
