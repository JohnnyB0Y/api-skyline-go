//  handlers_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/15.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// go test -run TestSendJSON -race -cpu 16

// Sample test to show how to test the execution of an internal endpoint.
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/testing/tests/example4/handlers"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	// 绑定路由路径
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint.
func TestSendJSON(t *testing.T) {
	url := "/sendjson" // API路径
	statusCode := 200  // 状态码

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		// 配置请求，传入操作方法 和 路径
		r := httptest.NewRequest("GET", url, nil)
		// 响应
		w := httptest.NewRecorder()
		// 模拟网络请求发送到服务器处理，这样就不需要运行服务器了。
		//
		http.DefaultServeMux.ServeHTTP(w, r)

		testID := 0
		t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, url, statusCode)
		{
			// 校验请求状态码
			if w.Code != 200 {
				t.Fatalf("\t%s\tTest %d:\tShould receive a status code of %d for the response. Received[%d].", failed, testID, statusCode, w.Code)
			}
			t.Logf("\t%s\tTest %d:\tShould receive a status code of %d for the response.", succeed, testID, statusCode)

			var u struct {
				Name  string
				Email string
			}

			// 校验JSON解析是否成功？
			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tTest %d:\tShould be able to decode the response.", failed, testID)
			}
			t.Logf("\t%s\tTest %d:\tShould be able to decode the response.", succeed, testID)

			// 校验响应的名字
			if u.Name == "Bill" {
				t.Logf("\t%s\tTest %d:\tShould have \"Bill\" for Name in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have \"Bill\" for Name in the response : %q", failed, testID, u.Name)
			}

			// 校验响应的邮件地址
			if u.Email == "bill@ardanlabs.com" {
				t.Logf("\t%s\tTest %d:\tShould have \"bill@ardanlabs.com\" for Email in the response.", succeed, testID)
			} else {
				t.Errorf("\t%s\tTest %d:\tShould have \"bill@ardanlabs.com\" for Email in the response : %q", failed, testID, u.Email)
			}
		}
	}
}
