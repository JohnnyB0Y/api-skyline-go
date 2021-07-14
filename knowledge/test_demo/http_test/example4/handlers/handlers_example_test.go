//  handlers_example_test.go
//
//
//  Created by JohnnyB0Y on 2021/07/15.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// go test -run ExampleSendJSON

// Sample to show how to write a basic example.
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// ExampleSendJSON provides a basic example example.
func ExampleSendJSON() {
	r := httptest.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(u)
	// Output:
	// {Bill bill@ardanlabs.com}
}
