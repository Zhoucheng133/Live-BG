package main

import "fmt"

var port string

func main() {
	fmt.Print("Input netPlayer ws port (Press Enter to use deafult port 9098): ")
	fmt.Scanln(&port)
	if port == "" {
		port = "9098"
		fmt.Println("➜ Use default port: 9098")
	}

}
