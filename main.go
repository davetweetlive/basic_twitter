package main

import (
	"firstapp/models"
	"firstapp/routes"
	"firstapp/utils"
	"net/http"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")

	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
