package services

import (
	"net/http"

	"github.com/Zawar-Ahmed10p/go_postgres/repositories"
	"github.com/Zawar-Ahmed10p/go_postgres/types"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		name := r.FormValue("name")
		email := r.FormValue("email")

		user := types.User{Name: name, Email: email}
		repositories.AddUser(user)

	}
}
