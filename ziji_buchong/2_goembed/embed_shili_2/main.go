package main

import (
	"embed"
	"net/http"
)

//参考文档
//https://blog.csdn.net/Naisu_kun/article/details/130722873#:~:text=%2F%2Fgo%3Aembed%20%E6%8C%87%E4%BB%A4%201%20%E8%A6%81%E5%B5%8C%E5%85%A5%E7%9A%84%E6%96%87%E4%BB%B6%E6%94%AF%E6%8C%81%E5%BD%93%E5%89%8D%E7%A8%8B%E5%BA%8F%20%28%20%2A.go%20%29%E6%89%80%E5%9C%A8%E7%9B%AE%E5%BD%95%E5%8F%8A%E5%AD%90%E7%9B%AE%E5%BD%95%EF%BC%9B%202,%E6%9D%A5%E5%8C%B9%E9%85%8D%E6%89%80%E6%9C%89%E6%96%87%E4%BB%B6%EF%BC%8C%E5%B9%B6%E4%B8%94%E4%BC%9A%E9%80%92%E5%BD%92%E5%AD%90%E7%9B%AE%E5%BD%95%E4%B8%AD%20.%20...%207%20%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%20%2F%2Fgo%3Aembed%20%2A%20%E6%9D%A5%E8%A1%A8%E7%A4%BA%E5%BD%93%E5%89%8D%E7%9B%AE%E5%BD%95%EF%BC%9B

// 地址后面的参数是文件路径
//
//go:embed index.html asset
//asset是文件夹路径(本例的当前目录是：embed_shili_2, asset意思是：embed_shili_2/asset路径)，是相对路径,
//asset是与main.go文件同级的文件夹。

//关于embed的路径：
//1. //go:embed 后也可以跟文件夹，默认情况下不会包含文件夹中. _ 开头的文件
//2. 可以使用 * 来匹配文件下所有文件（包括 . _ 开头的文件），但 * 不会递归子目录中 . _ 开头的文件
//3. 可以在文件夹名称前加上 all: 来匹配所有文件，并且会递归子目录中 . _ 开头的文件；
//4. 可以使用//go:embed * 来表示当前目录
//本例中的当前目录是：embed_shili_2
//。

var content embed.FS

func main() {
	// 使用embed.FS作为webServer静态文件服务目录
	fs := http.FileServer(http.FS(content))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

// go run main.go 后
// 在浏览器中，localhost:8080 可以查看到文字及效果
// localhost:8080/asset 可以查看该路径下的文件
// localhost:8080/asset/style  or localhost:8080/asset/style/index.css 均可查看相关文件或数据。
