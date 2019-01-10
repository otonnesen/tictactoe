package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT") // Get Heroku port

	if port == "" {
		InitLogger(os.Stdout, os.Stdout, os.Stdout, true)
		Info.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	} else {
		InitLogger(os.Stdout, os.Stdout, os.Stdout, false)
	}

	http.HandleFunc("/", LogRequest(Root))
	http.HandleFunc("/id/", LogRequest(Id))
	http.HandleFunc("/start/", LogRequest(Start))
	http.HandleFunc("/test/", LogRequest(Test))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	Info.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
