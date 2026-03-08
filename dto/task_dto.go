package dto

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required,oneof=todo in_progress done"`
	Priority    string `json:"priority" binding:"required,oneof=low medium high"`
	DueDate     string `json:"due_date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	AssigneeID  string `json:"assignee_id"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status" binding:"omitempty,oneof=todo in_progress done"`
	Priority    *string `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *string `json:"due_date" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	AssigneeID  *string `json:"assignee_id"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	DueDate     string `json:"due_date"`
	CreatorID   string `json:"creator_id"`
	AssigneeID  string `json:"assignee_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
