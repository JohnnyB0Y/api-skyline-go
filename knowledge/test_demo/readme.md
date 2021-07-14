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

- 性能测试：
  - 测试函数：func BenchmarkXXX(b *testing.B) {}

### 流程
1，当我们开始写代码的时候，可以在把源代码文件和测试文件放在同名的包里面，对未导出的API进行测试。
2，当我们需要编写导出的API的时候，可以把前面的测试丢弃，在不同的包里面对导出的API编写测试代码。

