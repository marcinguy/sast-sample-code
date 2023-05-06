package main

import (
  "log"
  "fmt"
  "net/http"
)

func serve() {
  http.HandleFunc("/clear_text_logging", func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    user := r.Form.Get("user")
    pw := r.Form.Get("password")

    log.Printf("Registering new user %s with password %s.\n", user, pw)
  })

  http.HandleFunc("/open_redirect", func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    http.Redirect(w, r, r.Form.Get("target"), 302)
  })

	http.HandleFunc("/xss", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")

    fmt.Fprintf(w, "%q is the username", username)
	})

  http.ListenAndServe(":80", nil)
}

