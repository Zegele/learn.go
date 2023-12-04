/*
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	UPLOAD_DIR   = "E:/Geek/src/learn.go/buchong/13wangluo/shili12_PhotoWeb/photoweb/uploads"
	TEMPLATE_DIR = "E:/Geek/src/learn.go/buchong/13wangluo/shili12_PhotoWeb/photoweb/view"
	PUBLIC_DIR   = "E:/Geek/src/learn.go/buchong/13wangluo/shili12_PhotoWeb/photoweb/public"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInforArr, err := ioutil.ReadDir(UPLOAD_DIR)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInforArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	//t, err := template.ParseFiles("E:/Geek/src/learn.go/buchong/13wangluo/shili12_PhotoWeb/photoweb/view/list.html")
	if err = renderHtml(w, TEMPLATE_DIR+"/list", locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderHtml(w http.ResponseWriter, templatePath string, locals map[string]interface{}) (err error) {
	t, err := template.ParseFiles(templatePath + ".html")
	if err != nil {
		return
	}
	err = t.Execute(w, locals)
	return err
}

//func main() {
//	http.HandleFunc("/", listHandler)
//	http.HandleFunc("/view", viewHandler)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err.Error())
//	}
//}

//加了闭包
//func main() {
//	http.HandleFunc("/", safeHanlder(listHandler))
//	http.HandleFunc("/view", safeHanlder(viewHandler))
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err.Error())
//	}
//}

//加了动静态分离
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", PUBLIC_DIR, 0)
	mux.HandleFunc("/", safeHanlder(listHandler))
	mux.HandleFunc("/view", safeHanlder(viewHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

var templates = make(map[string]interface{})

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
		fmt.Println("......", templates)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 闭包让程序更安全
func safeHanlder(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok { //recover().(error)是啥意思
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的50x错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Printf("WARN: panic in %v -  %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

// 动态，静态分离
const (
	ListDir = 0x0001
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}


 */