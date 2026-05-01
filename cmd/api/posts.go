package main

import (
	"net/http"

	"github.com/JayTailor45/go-social/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := ReadJSON(w, r, &payload); err != nil {
		WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := 1

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UsesrID: int64(userId),
		Tags:    payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := WriteJSON(w, http.StatusCreated, post); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
