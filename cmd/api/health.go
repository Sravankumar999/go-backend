package main

import (
	"go-dock/internal/store"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	post := &store.Post{Content: "Hello there"}
	app.store.Posts.Create(r.Context(), post)
}
