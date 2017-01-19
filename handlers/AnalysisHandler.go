package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/psionikangel/scribd-web/models"
)

//AnalysisHandler : Handles requests for analysis of a run
func AnalysisHandler(w http.ResponseWriter, r *http.Request, cfg models.Config) {
	if r.Method == "GET" {
		url := "http://" + cfg.Server + ":" + cfg.Port + "/analysis?runid=" + r.URL.Query().Get("runid")
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
		metas := new(models.MetadataList)
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(&metas)
		if err != nil {
			panic(err)
		}
		t, _ := template.ParseFiles("tmpl/home.html", "tmpl/analysis.html")
		t.Execute(w, metas)
	}
}
