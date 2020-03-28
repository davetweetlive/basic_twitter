package main

import (
	"go_web/controllers/demo_controller"
	"net/http"
)

func main() {
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", demo_controller.Index)
	http.HandleFunc("/demo/index", demo_controller.Index)

	http.ListenAndServe(":8080", nil)
}
