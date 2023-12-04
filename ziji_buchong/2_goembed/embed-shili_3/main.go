package main

import (
	"embed"
	"net/http"
)

// * 相当于把.go文件夹中的所有文件都嵌入进变量a中了
//
//go:embed *
var a embed.FS

//嵌入的安全性如何呢？

func main() {
	fs := http.FileServer(http.FS(a))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

//localhost:8080 就可以看到.go文件所在的文件夹中的所有文件
