package handlers

import (
	"net/http"
	"task_m/dto"
	"task_m/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
		return
	}

	// response, err := h.taskService.CreateTask(req)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
	// 	return
	// }
	// c.JSON(http.StatusCreated, response)
}
