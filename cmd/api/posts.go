package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/JayTailor45/go-social/internal/store"
	"github.com/go-chi/chi/v5"
)

type postKey string

const postCtxKey postKey = "post"

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags" validate:"required,max=1000"`
}
type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,required,max=100"`
	Content *string `json:"content" validate:"omitempty,required,max=1000"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
	}

	userId := 1

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  int64(userId),
		Tags:    payload.Tags,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r.Context())

	ctx := r.Context()

	comments, err := app.store.Comments.GetByPostID(ctx, post.ID)

	post.Comments = comments

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	err = app.store.Posts.Delete(ctx, int64(id))

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundError(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r.Context())

	var payload UpdatePostPayload
	if err := ReadJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if err := app.store.Posts.Update(r.Context(), post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, post); err != nil {
		app.internalServerError(w, r, err)
	}

}

// middleware to load post from database and add it to the request context
func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		post, err := app.store.Posts.GetById(ctx, int64(id))

		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundError(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, postCtxKey, post)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func getPostFromContext(ctx context.Context) *store.Post {
	post, _ := ctx.Value(postCtxKey).(*store.Post)
	return post
}
