package main

import (
	"log"
	"net/http"
	"time"

	"github.com/agentio/atiquette/api/com/atproto"
)

const AtprotoDatetimeLayout = "2006-01-02T15:04:05.999Z"

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "not found", http.StatusNotFound)
	}
	r.ParseForm()
	s := r.FormValue("status")
	did := r.Header.Get("user-did")
	status := map[string]interface{}{
		"status":    s,
		"createdAt": time.Now().UTC().Format(AtprotoDatetimeLayout),
	}
	authorizedClient := sessionClient.AuthorizedCopy(r)
	input := atproto.RepoCreateRecord_Input{
		Collection: "xyz.statusphere.status",
		Repo:       did,
		Record:     status,
	}
	out, err := atproto.RepoCreateRecord(r.Context(), authorizedClient, &input)
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	saveStatus(&Status{
		Uri:       out.Uri,
		AuthorDid: did,
		Status:    status["status"].(string),
		CreatedAt: status["createdAt"].(string),
		UpdatedAt: time.Now().UTC().Format(AtprotoDatetimeLayout),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
