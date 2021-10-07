package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed static
	res embed.FS
	pages = map[string]string{
		"/" : "static/index.gohtml",
	}
)

func main() {
	log.Println("booting simulation...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, ok := pages[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		tpl, err := template.ParseFS(res, page)
		if err != nil {
			log.Printf("page %s not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		if err := tpl.Execute(w, ""); err != nil {
			return
		}
	})

	http.FileServer(http.FS(res))
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
