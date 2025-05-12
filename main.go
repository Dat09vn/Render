package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Render!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := ":10000" // Port Render sử dụng khi không có $PORT
	if p := getenv("PORT", "10000"); p != "" {
		port = ":" + p
	}
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getenv(key, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return fallback
}
