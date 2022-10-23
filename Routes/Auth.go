package Routes

import (
	"foodways/Handlers"
	"foodways/Pkg/Mysql"
	"foodways/Repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := Repositories.RepositoryAuth(Mysql.DB)
	h := Handlers.HandlerAuth(userRepository)

	r.HandleFunc("/Register", h.Register).Methods("POST")
	r.HandleFunc("/Login", h.Login).Methods("POST")
}
