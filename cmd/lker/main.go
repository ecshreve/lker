package main

import (
	"log"
	"net/http"

	"golang.org/x/exp/slog"
)

func main() {
	slog.Info("---> main() - enter")
	defer slog.Info("<--- main() - exit")

	indexHandler := func(w http.ResponseWriter, _ *http.Request) {
		slog.Info("---> indexHandler() - enter")
		defer slog.Info("<--- indexHandler() - exit")
		w.Write([]byte("<h1>Hello World!</h1>"))
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8880", nil))
}
