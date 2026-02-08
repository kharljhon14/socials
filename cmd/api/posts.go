package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kharljhon14/socials/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=255"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags" validate:"required"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.statusBadRequestError(w, r, err)
		return
	}

	userID := uuid.New()

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  userID,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		app.statusInternalServerError(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, envelope{"data": post}, nil); err != nil {
		app.statusInternalServerError(w, r, err)
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postIDParam := chi.URLParam(r, "postID")

	postID, err := uuid.Parse(postIDParam)
	if err != nil {
		app.statusBadRequestError(w, r, err)
		return
	}

	ctx := r.Context()
	post, err := app.store.Posts.GetByID(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.statusNotFoundError(w, r, err)
		default:
			app.statusInternalServerError(w, r, err)

		}
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"data": post}, nil)
	if err != nil {
		app.statusInternalServerError(w, r, err)
	}
}
