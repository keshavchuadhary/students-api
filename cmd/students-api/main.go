package main

import (
	"net/http"

	"github.com/keshavchuadhary/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setup
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})
	//setup	server
	server := http.Server{
		Addr: cfg.Addr, Handler: router,
	}
	server.ListenAndServe()
}
