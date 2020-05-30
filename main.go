package main

import (
	"log"
	"net/http"

	"github.com/evanxzj/chitchat/config"
	"github.com/evanxzj/chitchat/routes"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	conf := config.LoadConfig()
	r := routes.NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(conf.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + conf.App.Address)
	err := http.ListenAndServe(conf.App.Address, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occurred starting HTTP listener at " + conf.App.Address)
		log.Println("Error: " + err.Error())
	}
}
