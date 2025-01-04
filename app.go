package main

import (
	"embed"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

var port string
var servicePort string

//go:embed web/dist/index.html
var indexHTML []byte

//go:embed web/dist/vite.svg
var viteSvg []byte

//go:embed web/dist/assets/*
var assets embed.FS

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(indexHTML)
}

func servePort(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(port))
}

func serveSvg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.WriteHeader(http.StatusOK)
	w.Write(viteSvg)
}
func serveAssets(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/assets/"):]
	file, err := assets.Open("web/dist/assets/" + path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Unable to stat file", http.StatusInternalServerError)
		return
	}

	// 动态获取文件的 MIME 类型
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
	w.Header().Set("Last-Modified", fileInfo.ModTime().UTC().Format(http.TimeFormat))
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Unable to send file", http.StatusInternalServerError)
	}
}

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

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/api/port", servePort)
	http.HandleFunc("/vite.svg", serveSvg)
	http.HandleFunc("/assets/", serveAssets)

	fmt.Printf("➜ Service starts at http://127.0.0.1:%s\n", servicePort)
	err := http.ListenAndServe(":"+servicePort, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
