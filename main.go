<<<<<<< HEAD
// filepath: /Users/mb/Projects/my-cli-app/main.go
=======
>>>>>>> f1f5527 (All changes: static assets, Docker, Go, and UI polish)
package main

import (
	"log"
	"net/http"
	"os"

	"jazzman/web"
)

func main() {
	http.HandleFunc("/", web.Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Jazzman web server running on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
