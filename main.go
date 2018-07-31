package main

import (
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	/*
	http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)
	http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	http.HandleFunc("/trash/", TrashTaskFunc)
	http.HandleFunc("/edit/", EditTaskFunc)
	http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	http.HandleFunc("/restore/", RestoreTaskFunc)
	http.HandleFunc("/add/", AddTaskFunc)
	http.HandleFunc("/update/", UpdateTaskFunc)
	http.HandleFunc("/search/", SearchTaskFunc)
	http.HandleFunc("/login", GetLogin)
	http.HandleFunc("/register", PostRegister)
	http.HandleFunc("/admin", HandleAdmin)
	http.HandleFunc("/add_user", PostAddUser)
	http.HandleFunc("/change", PostChange)
	http.HandleFunc("/logout", HandleLogout)*/
	http.HandleFunc("/", ShowAllTasksFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {

	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}

	w.Write([]byte(message))
}