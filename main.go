package main

import "net/http"

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", movieIndex)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)


	mux.HandleFunc("/movie/:id", Show)
	mux.HandleFunc("/movie/create", Create)
	mux.HandleFunc("/movie/:id/update", Update)
	mux.HandleFunc("/movie/:id/delete", Delete)


	server := &http.Server {
		Addr: 			"0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
