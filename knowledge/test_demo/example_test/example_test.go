package exampletest

func ExampleSayHello() {
	SayHello()
	// Output: Hello World!
}

func ExampleSayYes() {
	SayYes()
	// Output:
	// Yes.
	// Yes.
}

func ExampleSayNo() {
	SayNo()
	// Unordered Output:
	// No
	// No.
	// No..
	// No...
}

/**
执行测试
go test -v -cover

检测单行输出格式为 // Output: <期望字符串>
检测多行输出格式为 // Output: <期望字符串> <期望字符串>
检测无序输出格式为 // Unordered Output: <期望字符串> <期望字符串> etc.
*/
