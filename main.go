package main

import (
	"ProjectAuthenticate/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Exicted URL: /login, /logout, /del, /refresh"))

	})
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("POST")
	r.HandleFunc("/del", controllers.DelAll).Methods("POST")
	r.HandleFunc("/refresh", controllers.Refresh).Methods("POST")

	_ = http.ListenAndServe(os.Getenv("PORT"), r)
}
