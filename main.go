package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	log.Println("hello world")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /scramble", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("java", "-jar", "tnoodle.jar", "scramble")
		out, err := cmd.Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(out)
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("listening on :8080")
	log.Fatal(server.ListenAndServe())

}
