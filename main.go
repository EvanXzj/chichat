package main

import (
	"log"
	"net/http"

	"github.com/evanxzj/chitchat/routes"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	r := routes.NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
