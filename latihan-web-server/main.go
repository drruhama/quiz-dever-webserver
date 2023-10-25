package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called...")
	hello := "Hello World"
	w.Write([]byte(hello))
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About called...")
	about := "Ini adalah halaman About"
	w.Write([]byte(about))
}

func handeRoute() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
}

func main() {
	port := ":9999"
	startServer(port)

}

func startServer(port string) {
	handeRoute()
	fmt.Println("Server running at port", port)
	http.ListenAndServe(port, nil)
}
