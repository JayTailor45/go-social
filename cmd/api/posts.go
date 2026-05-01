package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/JayTailor45/go-social/internal/store"
	"github.com/go-chi/chi/v5"
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

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetById(ctx, int64(id))

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			WriteJsonError(w, http.StatusNotFound, err.Error())
		default:
			WriteJsonError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := WriteJSON(w, http.StatusOK, post); err != nil {
		WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
