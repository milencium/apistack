package http
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/milenicum/internal/comment"
)
// Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
	Service *comment.Service
}
// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service)*Handler{
	return &Handler{
		Service: service
	}
}
// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes(){
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "I am alive server")
	})
	//Adding new handle functions
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
}
// GetComment - retreve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UNIT from ID")
	}
	//tu pozivamo već kreiranu pravu service funkciju
	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error retrieveing comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)

}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request){
	//tu pozivamo već kreiranu pravu service funkciju
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve all comments")
	}
	fmt.Fprintf(w, "%v", comments)
}

// PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request){
	//tu pozivamo već kreiranu pravu service funkciju
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/", 
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprinf(w, "%v", comment)
}

// UpdateComment - updates a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request){
	//tu pozivamo već kreiranu pravu service funkciju
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	fmt.Fprintf(w, "%v", comment)
}
// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to pass UINT from ID")
	}
	//tu pozivamo već kreiranu pravu service funkciju
	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment by comment ID")
	}
	fmt.Fprinf(w, "Successfully deleted comment")
}