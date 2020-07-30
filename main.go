package main

import (
	"basic_twitter/models"
	"basic_twitter/routes"
	"basic_twitter/utils"
	"net/http"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")

	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
