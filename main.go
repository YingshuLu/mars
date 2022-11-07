package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yingshulu/mars/config"
	"github.com/yingshulu/mars/rest"
)

//go:embed misc/*
var staticFS embed.FS

func restfulRoute(r *gin.Engine) {
	group := r.Group("/rest")
	group.POST("/register", rest.Register)
	group.POST("/login", rest.Login)
}

func staticFSRoute(r *gin.Engine) {
	misc, _ := fs.Sub(staticFS, "misc")

	css, _ := fs.Sub(misc, "css")
	r.StaticFS("/css", http.FS(css))
	js, _ := fs.Sub(misc, "script")
	r.StaticFS("/script", http.FS(js))

	video, _ := fs.Sub(misc, "video")
	videoGroup := r.Group("/video", rest.Auth)
	videoGroup.StaticFS("/", http.FS(video))

	html, _ := fs.Sub(misc, "html")
	htmlGroup := r.Group("/html", rest.Auth)
	htmlGroup.StaticFS("/", http.FS(html))

	r.StaticFileFS("/index.html", "login.html", http.FS(html))
	r.StaticFileFS("/register.html", "register.html", http.FS(html))
	r.StaticFileFS("/login.html", "login.html", http.FS(html))
}

func route(r *gin.Engine) *gin.Engine {
	staticFSRoute(r)
	restfulRoute(r)
	return r
}

func run(r *gin.Engine) *gin.Engine {
	cnf := config.Global()
	port := fmt.Sprintf(":%v", cnf.Host.Port)
	if cnf.Host.UseSSL == 0 {
		r.Run(port)
	} else {
		r.RunTLS(port, cnf.Host.CertPath, cnf.Host.KeyPath)
	}
	return r
}

func main() {
	run(route(gin.Default()))
}
