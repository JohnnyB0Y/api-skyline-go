# 测试驱动开发

### 约定
测试文件以 _test 结尾：xxx_test.go

- 单元测试：t *testing.T
  - 测试函数：func TestXXX(t *testing.T) {}
    - t.Error 测试失败并记录信息
    - t.Log 测试记录信息
    - t.Fatal 测试失败，记录信息，并结束测试


- 示例测试：
  - 测试函数：func ExampleXXX() {}
  - 既可以出现在文档里，也可以用来执行测试
    - [打开本地文档：](https://learnku.com/articles/56426)~/go/bin/godoc -http=:6060 -play

- 基准测试：
  - 测试函数：func BenchmarkXXX(b *testing.B) {}
  - 测试参数
    - go test -run none -bench . -benchtime 3s -benchmem
      - -run none 让 go test 匹配不到所有测试文件（我们没有以 none 做测试函数名，所以才匹配不到）
      - -bench . 表示对所有的 benchmark函数 做测评
      - -benchtime 3s 设置测试时间 3秒（默认是1s）
      - -benchmem 输出内存分配情况
      - -memprofile m.out 生成内存分析报告
  - 注意事项
    - 进行基准测试前，电脑性能需要保证最优；
    - 多个基准测试，可能会互相影响测试结果；（运行下一次测试时，可能同时也在清理上一次测试分配的内存或资源）


### [命令行](https://hyper0x.github.io/go_command_tutorial/#/0.7)
- 开始测试
  - go test

参数
  - -v：表示打印详细信息
  - -cover：打印测试覆盖率
  - -run：匹配测试
  
匹配测试，匹配 -run 后面的 Down，系统会找到所有 匹配到 Down 字符的测试函数开始测试。
  - go test -run Down

覆盖度详情报告
  - 生成测试覆盖度详情报告
    - go test -coverprofile c.out
  - 在网页里查看覆盖度详情报告
    - go tool cover -html c.out

内存分析报告
  - 生成内存分析报告
    - go test -run none -bench . -benchtime 3s -benchmem -memprofile m.out
      - -memprofile m.out 生成内存分析报告
    - go test -gcflags "-m -m" -run none -bench . -benchtime 3s -benchmem -memprofile m.out
      - -gcflags "-m -m" 让编译器在运行测试之前，生成一份逸出分析报告
  - 查看分析报告
    - go tool pprof -alloc_space m.out
  - pprof 命令
    - list algOne 找出调用algOne函数的内存分配信息

CPU分析报告
  - 生成CPU分析报告
    - go test -run none -bench . -benchtime 3s -benchmem -cpuprofile c.out
  - 查看分析报告
    - go tool pprof c.out
  - 在网页中查看分析报告
    - go tool pprof -http :3000 c.out
  - pprof 命令
    - list algOne 找出调用algOne函数的CPU处理信息

负载追踪报告
  - 让系统每秒生成一份调度器追踪信息
    - > GODEBUG=schedtrace=1000 ./project > /dev/null
      - > schedtrace=1000 设置1000毫秒（1秒）追踪一次；
      - > \> /dev/null 默认错误信息是输出到错误端，这里是把错误信息从定向到 /dev/null 文件夹里；
  - 让系统每秒生成一份垃圾收集追踪信息 (垃圾收集时间最好不要超过100微妙，两次垃圾收集大概在1到2毫秒)
    - > GODEBUG=gctrace=1 ./project > /dev/null
  - 运行负载请求
    - > hey -m POST -c 100 -n 10000 "http://localhost:5000/search>term=trump&cnn=on&bbc=on&nyt=on"
      - > -m POST 请求方法
      - > -c 100 启动100条连接
      - > -n 10000 发起请求 10000 次
      - > "xxxx" 请求的 URL


### 流程
1，当我们开始写代码的时候，可以在把源代码文件和测试文件放在同名的包里面，对未导出的API进行测试。
2，当我们需要编写导出的API的时候，可以把前面的测试丢弃，在不同的包里面对导出的API编写测试代码。

