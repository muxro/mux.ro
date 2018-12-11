package main // import "github.com/muxro/mux.ro"

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gobuffalo/packr"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.Recoverer)

	staticBox := packr.NewBox("./static")
	r.Mount("/", http.FileServer(staticBox))

	http.ListenAndServe(":"+port, r)
}
