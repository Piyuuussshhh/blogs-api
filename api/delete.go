package api

import (
	"blog-api/db"
	"fmt"
	"net/http"
)

func deleteBlog(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	if id == "" {
		http.Error(w, "[ERROR] Id is not provided.", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteBlogById(req.Context(), id); err != nil {
		http.Error(w, fmt.Sprintf("[ERROR] %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}