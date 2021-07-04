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

```Go
// 映射表的数据结构
type custom_map struct {
  count int // 元素个数
  B uint8 // bucket 数组大小
  buckets unsafe.Pointer // 装载 bucket 的数组，长度为2的B次方；（B是单个bucket数组的大小）
  oldbuckets unsafe.Pointer // 老旧的 buckets，扩容迁移元素时使用。
  ... // 其他一些
}

// bucket 的数据结构
type custom_bucket struct {
  tophash [8]uint8 // 存储hash值的高8位的数组
  data []bype // 键值结构数据，按 key1, key2 ... key8, val1, val2 ... val8 存储
  overflow *custom_bucket // 当前bucket溢出时，指向下一个bucket的指针
}

```
