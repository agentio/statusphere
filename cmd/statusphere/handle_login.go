package main

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("login").Parse(login_template)
	if err != nil {
		return
	}
	err = t.ExecuteTemplate(w, "login", map[string]any{
		"Prefix": prefix,
	})
	if err != nil {
		return
	}
}

const login_template = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Log in</title>
<link rel="stylesheet" href="/public/styles.css">
</head>
<body>
<div id="root">
<div id="header">
<h1>Statusphere</h1>
<p>Set your status on the Atmosphere.</p>
</div>
<div class="container">
<form action="/{{.Prefix}}/login" method="post" class="login-form">
<input type="text" name="handle" placeholder="Enter your handle (eg alice.bsky.social)" required="">
<button type="submit">Log in</button>
</form>
<div class="signup-cta">
Don't have an account on the Atmosphere?
<a href="https://bsky.app/">Sign up for Bluesky</a> to create one now!
</div>
</div>
</div>
</body></html>
`
