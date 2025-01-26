package ctrl

import (
	"app/internal/pkg/msvc.tracker/post/svc"
)

// PostController handles the operations related to domain items.
type PostController struct {
	trackerService *svc.PostService // Service for fetching somethings.
}

// NewPostController creates a new PostController with the provided service.
func NewPostController(service *svc.PostService) *PostController {
	return &PostController{trackerService: service}
}
