package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var inp domain.CreateTaskRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	task, err := h.service.StartTask(c, inp)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error creating task: %v", err))
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) StopTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id param: %v", err))
		return
	}

	var inp domain.StopTaskRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	err = h.service.StopTask(c, taskId, inp.UserId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully stopped"})
}

func (h *Handler) GetTaskProgressByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request id: %v", err))
		return
	}

	tasks, err := h.service.GetTaskProgressByUserId(c, id)
	if err != nil {
		if err == domain.ErrNoTask {
			errorResponse(c, http.StatusBadRequest, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	if len(tasks) == 0 {
		errorResponse(c, http.StatusNotFound, fmt.Errorf("no tasks found for user ID %d", id))
		return
	}

	c.JSON(http.StatusOK, tasks)
}
