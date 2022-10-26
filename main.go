package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed misc/*
var staticFS embed.FS

func main() {
	route := gin.Default()
	fe, _ := fs.Sub(staticFS, "misc")
	route.StaticFS("/static", http.FS(fe))

	route.Run(":8080")
}
