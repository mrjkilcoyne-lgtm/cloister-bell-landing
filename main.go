// Cloister Bell landing — single-binary htmx site.
//
// The thesis of the artefact (a thought-decoder for sensor-fused biosignals
// from off-the-shelf earbuds + watch + ring) is *anti-SPA*: latency matters,
// the wire format matters, and rendering happens where the data lives.
// This site embodies that thesis. One binary. Server-rendered HTML.
// htmx for the small interactive bits.
package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*.html
var templateFS embed.FS

//go:embed static/*
var staticFS embed.FS

type pageData struct {
	Title       string
	Description string
	Year        int
}

func main() {
	addr := flag.String("addr", envOr("ADDR", ":8080"), "listen address")
	flag.Parse()

	tmpl, err := template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		log.Fatalf("parse templates: %v", err)
	}

	staticSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatalf("static sub: %v", err)
	}

	mux := http.NewServeMux()

	// Static assets — long cache, content-hashed in prod (PLACEHOLDER: hash step).
	mux.Handle("/static/", http.StripPrefix("/static/", cacheControl(http.FileServer(http.FS(staticSub)), "public, max-age=3600")))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFound(w, tmpl)
			return
		}
		render(w, tmpl, "index.html", pageData{
			Title:       "Cloister Bell — thought decoder for the sensor-fusion era",
			Description: "Decoding intent from earbud, watch and ring biosignals — without a training run.",
			Year:        time.Now().UTC().Year(),
		})
	})

	mux.HandleFunc("/thesis", func(w http.ResponseWriter, r *http.Request) {
		render(w, tmpl, "thesis.html", pageData{
			Title:       "Thesis — Cloister Bell",
			Description: "Why thought-decoding works without ML training, given the right grammar bridges.",
			Year:        time.Now().UTC().Year(),
		})
	})

	mux.HandleFunc("/hardware", func(w http.ResponseWriter, r *http.Request) {
		render(w, tmpl, "hardware.html", pageData{
			Title:       "Hardware — Cloister Bell",
			Description: "Off-the-shelf sensors that make this practical: earbuds, watch, ring.",
			Year:        time.Now().UTC().Year(),
		})
	})

	// htmx partial — newsletter signup confirmation.
	mux.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// PLACEHOLDER: actual subscribe handler. For now, echo.
		_ = r.ParseForm()
		email := r.FormValue("email")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<p class="ok" role="status">Heard. We will write to <strong>%s</strong> when there is something to say.</p>`, template.HTMLEscapeString(email))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	log.Printf("cloister-bell-landing listening on %s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

func render(w http.ResponseWriter, tmpl *template.Template, page string, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	if err := tmpl.ExecuteTemplate(w, page, data); err != nil {
		log.Printf("render %s: %v", page, err)
		http.Error(w, "template error", http.StatusInternalServerError)
	}
}

func notFound(w http.ResponseWriter, tmpl *template.Template) {
	w.WriteHeader(http.StatusNotFound)
	_ = tmpl.ExecuteTemplate(w, "404.html", pageData{
		Title: "Not found — Cloister Bell",
		Year:  time.Now().UTC().Year(),
	})
}

func cacheControl(next http.Handler, value string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", value)
		next.ServeHTTP(w, r)
	})
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
