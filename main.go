package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ua := r.UserAgent()
		logger.Info("request", "user-agent", ua, "method", r.Method, "path", r.URL.Path, "remote-addr", r.RemoteAddr)
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Println("server is running on port 8080...")

	http.ListenAndServe(":8080", nil)
}
