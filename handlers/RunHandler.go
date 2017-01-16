package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/psionikangel/scribd-web/models"
)

//RunHandler : Handles requests for run information
func RunHandler(w http.ResponseWriter, r *http.Request, cfg models.Config) {
	if r.Method == "GET" {
		url := "http://" + cfg.Server + ":" + cfg.Port + "/run"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		runs := new(models.Runlist)
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&runs)
		if err != nil {
			panic(err)
		}
		t, _ := template.ParseFiles("tmpl/home.html", "tmpl/runs.html")
		t.Execute(w, runs)
	}
}
