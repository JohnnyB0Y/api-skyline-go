package exampletest

import "fmt"

func SayHello() {
	fmt.Println("Hello World!")
}

func SayYes() {
	fmt.Println("Yes.")
	fmt.Println("Yes.")
}

func SayNo() {

	nos := make(map[int]string, 4)
	nos[1] = "No."
	nos[2] = "No.."
	nos[3] = "No..."
	nos[0] = "No"

	for i := 0; i < 4; i++ {
		fmt.Println(nos[i])
	}
}
