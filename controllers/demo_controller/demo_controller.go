package demo_controller

import (
	"html/template"
	"net/http"
)

type Data struct {
	Age      int
	Username string
	// Product  entities.Product
}

var sherrifTmpl = template.New("test").Delims("{[{", "}]}")

func Index(response http.ResponseWriter, request *http.Request) {
	// tmplt, _ := template.ParseFiles("/home/dave/Go_workstation/src/go_web/views/demo_controller/index.html")
	data := Data{
		Age:      27,
		Username: "davetweetlive",
		// Product: entities.Product{
		// 	Id:       "P01",
		// 	Name:     "Name1",
		// 	Photo:    "thumb1.jpg",
		// 	Price:    5,
		// 	Quantity: 3,
		// },
	}
	template.Must(sherrifTmpl.ParseFiles("./views/demo_controller/index.html")).Execute(response, data)
	// tmplt.Execute(response, data)
}
