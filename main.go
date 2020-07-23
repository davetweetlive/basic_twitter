package main

import (
	"context"
	"net/http"
	"text/template"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
var templates *template.Template

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/test", testGetHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	comments, err := client.LRange(ctx, "comments", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ctx := context.TODO()
	comment := r.PostForm.Get("comment")
	client.LPush(ctx, "comments", comment)
	http.Redirect(w, r, "/", 302)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
func testGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untype, ok := session.Values["username"]

	if !ok {
		return
	}
	username, ok := untype.(string)
	if !ok {
		return
	}
	w.Write([]byte(username))
}
