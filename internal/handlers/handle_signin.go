package handlers

import (
	"html/template"
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("signin").Parse(signin_template)
	if err != nil {
		return
	}
	err = t.ExecuteTemplate(w, "signin", map[string]any{})
	if err != nil {
		return
	}
}

const signin_template = `
<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>Sign in</title>
<link rel="stylesheet" href="/public/styles.css">
</head>
<body>
<div id="root">
<div id="header">
<h1>Statusphere</h1>
<p>Set your status on the Atmosphere.</p>
</div>
<div class="container">
<form action="/@/signin" method="post" class="signin-form">
<input type="text" name="handle" placeholder="Enter your handle (eg alice.bsky.social)" required="">
<button type="submit">Sign in</button>
</form>
<div class="signup-cta">
Don't have an account on the Atmosphere?
<a href="https://bsky.app/">Sign up for Bluesky</a> to create one now!
</div>
</div>
</div>
</body></html>
`
