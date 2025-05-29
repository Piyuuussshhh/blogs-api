package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"blog-api/db"
	"blog-api/models"
)

func addBlog(w http.ResponseWriter, req *http.Request) {
	var jsonBody struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Category string   `json:"category"`
		Tags     []string `json:"tags"`
	}
	err := json.NewDecoder(req.Body).Decode(&jsonBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	
	newBlog := models.New(
		jsonBody.Title,
		jsonBody.Content,
		jsonBody.Category,
		jsonBody.Tags,
	)

	
	if err = db.InsertBlog(req.Context(), newBlog); err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newBlog); err != nil {
		http.Error(w, "[ERROR] Could not encode newly created blog to JSON", http.StatusInternalServerError)
		return
	}
}
