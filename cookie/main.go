package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"testing"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "hello world!")
	return
}


