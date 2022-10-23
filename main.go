package main

import (
	"fmt"
	"foodways/Database"
	"foodways/Pkg/Mysql"
	"foodways/Routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Mysql.DatabaseInit()

	Database.RunMigration()

	r := mux.NewRouter()

	Routes.RounteInit(r.PathPrefix("/api/v1/").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Localhost run on port 5000")
	http.ListenAndServe("localhost:5000", r)
}
