//  handlers.go
//
//
//  Created by JohnnyB0Y on 2021/07/15.
//  Copyright © 2021 JohnnyB0Y. All rights reserved.

// Package handlers provides the endpoints for the web service.
package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes sets the routes for the web service.
func Routes() {
	// 绑定API路径和函数
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple JSON document.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanlabs.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
