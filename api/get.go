package api

import (
	"blog-api/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func getBlog(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	if id == "" {
		http.Error(w, "[ERROR] Id is not provided.", http.StatusInternalServerError)
		return
	}

	blog, err := db.GetBlogByID(req.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(blog); err != nil {
		http.Error(w, "[ERROR] Could not encode the blog to JSON", http.StatusInternalServerError)
		return
	}
}

func getBlogBySearch(w http.ResponseWriter, req *http.Request) {
	term := req.URL.Query().Get("search")
	if term == "" {
		http.Error(w, "[ERROR] Term is not provided.", http.StatusInternalServerError)
		return
	}

	matchingBlogs, err := db.GetBlogsBySearch(req.Context(), term)
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(matchingBlogs); err != nil {
		http.Error(w, "[ERROR] Could not encode the blogs to JSON", http.StatusInternalServerError)
		return
	}
}

func getAllBlogs(w http.ResponseWriter, req *http.Request) {
	blogs, err := db.GetAllBlogs(req.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(blogs); err != nil {
		http.Error(w, "[ERROR] Could not encode the blogs to JSON", http.StatusInternalServerError)
		return
	}
}