package api

import (
	"blog-api/db"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func updateBlog(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	if id == "" {
		http.Error(w, "[ERROR] Id is not provided.", http.StatusInternalServerError)
		return
	}

	toBeUpdatedBlog, err := db.GetBlogByID(req.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v.\n", err.Error()), http.StatusInternalServerError)
		return
	}

	var jsonBody struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}
	if err := json.NewDecoder(req.Body).Decode(&jsonBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toBeUpdatedBlog.Title = jsonBody.Title
	toBeUpdatedBlog.Content = jsonBody.Content
	toBeUpdatedBlog.Category = jsonBody.Category
	toBeUpdatedBlog.Tags = jsonBody.Tags
	// Very important: set modifiedAt to current time.
	now := time.Now()
	toBeUpdatedBlog.ModifiedAt = &now

	if err := db.UpdateBlog(req.Context(), id, *toBeUpdatedBlog); err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(*toBeUpdatedBlog); err != nil {
		http.Error(w, "[ERROR] Could not encode updated blog to JSON", http.StatusInternalServerError)
		return
	}
}

/*
	1. Get the original blog.
	2. Update the fields and modifiedAt.
	3. Replace original with updated blog.
*/