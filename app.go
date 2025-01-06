package main

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var port string
var servicePort string

//go:embed web/dist/*
var staticFiles embed.FS

func main() {
	fmt.Print("Input netPlayer ws port (Press Enter to use deafult port 9098): ")
	fmt.Scanln(&port)
	if port == "" {
		port = "9098"
		fmt.Println("➜ Use default port: 9098")
	}
	fmt.Print("Input this service port (Press Enter to use default port 5000): ")
	fmt.Scanln(&servicePort)
	if servicePort == "" {
		servicePort = "5000"
		fmt.Println("➜ Use default service port: 5000")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/*path", func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/port") {
			c.String(200, port)
		} else {
			filePath := c.Request.URL.Path
			if filePath == "" || filePath == "/" {
				filePath = "web/dist/index.html"
			} else {
				filePath = "web/dist" + filePath
			}
			fmt.Println(filePath)
			file, err := staticFiles.Open(filePath)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("File not found: %s", filePath)})
				return
			}
			defer file.Close()
			ext := filepath.Ext(filePath)
			contentType := mime.TypeByExtension(ext)
			c.Header("Content-Type", contentType)
			c.DataFromReader(http.StatusOK, -1, filePath, file, map[string]string{})
		}
	})

	fmt.Printf("➜ Service starts at http://127.0.0.1:%s\n", servicePort)
	r.Run(fmt.Sprintf(":%s", servicePort))

}
