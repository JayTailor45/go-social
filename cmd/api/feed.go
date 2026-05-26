package main

import (
	"log"
	"net/http"

	"github.com/JayTailor45/go-social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// pagination, filters
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	// TODO: get user id from context
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(104), fq)
	if err != nil {
		app.internalServerError(w, r, err)
		log.Print(err.Error())
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
		log.Print(err.Error())
	}

}
