package ctrl

import "app/internal/pkg/msvc.tracker/record/svc"

// RecordController handles the operations related to domain items.
type RecordController struct {
	trackerService *svc.RecordService // Service for fetching somethings.
}

// NewRecordController creates a new RecordController with the provided service.
func NewRecordController(service *svc.RecordService) *RecordController {
	return &RecordController{trackerService: service}
}
