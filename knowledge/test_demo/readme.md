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

### 流程
1，当我们开始写代码的时候，可以在把源代码文件和测试文件放在同名的包里面，对未导出的API进行测试。
2，当我们需要编写导出的API的时候，可以把前面的测试丢弃，在不同的包里面对导出的API编写测试代码。

