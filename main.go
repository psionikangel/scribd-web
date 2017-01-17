package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/psionikangel/scribd-web/handlers"
	"github.com/psionikangel/scribd-web/models"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.HandleFunc("/metadata", makeHandler(handlers.MetaHandler))
	http.HandleFunc("/run", makeHandler(handlers.RunHandler))
	http.ListenAndServe(":8181", nil)
}

func loadConfig() *models.Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	cfg := new(models.Config)
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, models.Config)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := loadConfig()
		fn(w, r, *cfg)
	}
}
