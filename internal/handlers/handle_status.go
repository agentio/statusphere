package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/agentio/slink/pkg/client"
	"github.com/agentio/statusphere/gen/xrpc"
	"github.com/agentio/statusphere/internal/storage"
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
	c := client.NewClientWithOptions(client.ClientOptions{
		Host: sessionproxy(),
	})
	c.SetSessionHeaders(r)

	input := xrpc.ComATProtoRepoCreateRecord_Input{
		Collection: "xyz.statusphere.status",
		Repo:       did,
		Record:     status,
	}
	out, err := xrpc.ComATProtoRepoCreateRecord(r.Context(), c, &input)
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	storage.SaveStatus(&storage.Status{
		Uri:       out.Uri,
		AuthorDid: did,
		Status:    status["status"].(string),
		CreatedAt: status["createdAt"].(string),
		UpdatedAt: time.Now().UTC().Format(AtprotoDatetimeLayout),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
