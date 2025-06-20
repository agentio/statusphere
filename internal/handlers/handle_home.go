package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/agentio/atiquette/api/app/bsky"
	"github.com/agentio/atiquette/api/com/atproto"
	"github.com/agentio/statusphere/api/xyz/statusphere"
	"github.com/agentio/statusphere/internal/clients"
	"github.com/agentio/statusphere/internal/storage"
)

var handles = make(map[string]string)

func getHandle(did string) string {
	h, ok := handles[did]
	if ok {
		return h
	}
	result, err := bsky.ActorGetProfile(context.TODO(), clients.AnonymousClient, did)
	if err != nil {
		log.Printf("%s", err)
		return ""
	}
	h = result.Handle
	handles[did] = h
	return h
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").Funcs(template.FuncMap{
		"handle": func(s *storage.Status) string {
			return getHandle(s.AuthorDid)
		},
	}).Parse(home_template)
	if err != nil {
		return
	}

	// When a user is signed-in, IO puts their did in this header.
	did := r.Header.Get("user-did")

	statuses, _ := storage.ListStatus()
	var status string

	var dn string
	if did != "" {
		dn = did
		authorizedClient := clients.SessionClient.AuthorizedCopy(r) // configure the client with headers sent with this request
		// get user profile
		{
			out, err := bsky.ActorGetProfile(r.Context(), authorizedClient, did)
			if err == nil {
				if out.DisplayName != nil {
					dn = *out.DisplayName
				}
			} else {
				log.Printf("%+v", err)
			}
		}
		// get user status
		{
			out, err := atproto.RepoListRecords(r.Context(), authorizedClient,
				"xyz.statusphere.status", "", 1, did, false)
			if err == nil {
				if len(out.Records) > 0 {
					b, _ := json.Marshal(out.Records[0].Value)
					var s statusphere.Status
					err = json.Unmarshal(b, &s)
					if err != nil {
						fmt.Println(err)
					}
					status = s.Status
				}
			}
		}

	}

	err = t.ExecuteTemplate(w, "home", map[string]any{
		"Did":      did,
		"Name":     dn,
		"Buttons":  buttons,
		"Statuses": statuses,
		"Status":   status,
	})
	if err != nil {
		return
	}
}

var buttons = []string{
	"👍",
	"👎",
	"💙",
	"🥹",
	"😧",
	"😤",
	"🙃",
	"😉",
	"😎",
	"🤓",
	"🤨",
	"🥳",
	"😭",
	"😤",
	"🤯",
	"🫡",
	"💀",
	"✊",
	"🤘",
	"👀",
	"🧠",
	"👩‍💻",
	"🧑‍💻",
	"🥷",
	"🧌",
	"🦋",
	"🚀",
	"🌉",
}

const home_template = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Home</title>
<link rel="stylesheet" href="./public/styles.css">
</head>
<body>
<div id="root">
<div class="error"></div>
<div id="header">
<h1>Statusphere</h1>
<p>Set your status on the Atmosphere.</p>
</div>
<div class="container">
<div class="card">
<div class="session-form">
{{ if .Name }}
<div>Hi, <strong>{{ .Name }}</strong>. What's your status today?</div>
{{ else }}
<div><a href="/signin">Sign in</a> to set your status!</div>
{{ end }}
<div>
{{ if .Name }}
<a href="/@/signout" class="button">Sign out</a>
{{ else }}
<a href="/signin" class="button">Sign in</a>
{{ end }}
</div>
</div>
</div>
<form action="/status" method="post" class="status-options">
{{range $index, $element := .Buttons}}
<button name="status" class={{if eq $element $.Status}}"status-option selected"{{else}}"status-option"{{end}} value="{{ $element }}">{{ $element }}</button>
{{end}}
</form>
{{range $index, $status := .Statuses}}
<div class="status-line {{ if eq $index 0 }}no-line{{ end }}">
<div>
<div class="status">{{$status.Status}}</div>
</div>
<div class="desc">
<a class="author" href="https://bsky.app/profile/{{handle $status}}">@{{handle $status}}</a>
was feeling {{$status.Status}} on {{$status.Created}}
</div>
</div>
{{end}}

</body></html>
`
