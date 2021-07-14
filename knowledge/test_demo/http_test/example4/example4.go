//  example4.go
//
//
//  Created by JohnnyB0Y on 2021/07/15.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// 仿制服务端
// Sample program that implements a simple web service.
package main

import (
	"log"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/go/testing/tests/example4/handlers"
)

func main() {
	handlers.Routes()

	// 提供服务
	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)

	// 先编译成 example4 程序
	// ./example4 运行程序
	// 在浏览器 访问 http://localhost:4000/sendjson

}
