package api

import (
	"net/http"
)

func Serve() error {
	/// Add a post.
	http.HandleFunc("POST /posts", addBlog)
	/// Get all posts OR search by term.
	http.HandleFunc("GET /posts", func(w http.ResponseWriter, req *http.Request) {
		// use req.URL.Query().Get("smth") if the path is like /<main>?smth=""
		// use req.PathValue("smth") if the path is like /<main>/smth
		search := req.URL.Query().Get("search")
		if search == "" {
			getAllBlogs(w, req)
		} else {
			/// Case insensitive search btw.
			getBlogBySearch(w, req)
		}
	})
	/// Get a particular post with id = {id}.
	http.HandleFunc("GET /posts/{id}", getBlog)
	/// Update a particular post with id = {id}.
	http.HandleFunc("PUT /posts/{id}", updateBlog)
	/// Delete a particular post with id = {id}.
	http.HandleFunc("DELETE /posts/{id}", deleteBlog)

	return http.ListenAndServe(":8080", nil)
}