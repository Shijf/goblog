/*
 * @Github: https://github.com/shijf
 * @Author: shijf
 * @Date: 2020-12-26 08:42:56
 * @LastEditTime: 2020-12-26 08:42:56
 * @LastEditors: shijf
 * @FilePath: /golang.learnku.com/main.go
 * @Description:
 */
package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
