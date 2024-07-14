package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type user struct {
	Name  string `json:"user_name"`
	Email string `json:"user_email"`
}

var users = []user{}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(routeHandler))
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		if r.Method == http.MethodGet {

			data, _ := json.Marshal(users)

			w.WriteHeader(http.StatusAccepted)
			w.Write(data)
		} else if r.Method == http.MethodPost {
			newUser := user{}

			data, _ := io.ReadAll(r.Body)
			json.Unmarshal(data, &newUser)

			for _, v := range users {
				if v.Email == newUser.Email {
					w.WriteHeader(http.StatusBadRequest)

					return
				}
			}

			users = append(users, newUser)

			w.WriteHeader(http.StatusCreated)

			d, _ := json.Marshal(newUser)

			w.Write(d)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
