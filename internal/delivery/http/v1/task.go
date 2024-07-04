package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/pkg"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Creates a new task by taking a name and description
// @Tags task
// @Accept  json
// @Produce  json
// @Param   input  body      domain.CreateTaskRequest  true  "Task Creation Data"
// @Success 201  {object} entity.Task
// @Failure 400  {object}  Response
// @Failure 500 {object} Response
// @Router /tasks [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var inp domain.CreateTaskRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if !pkg.ValidateStrings(inp.Name, inp.Description) {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("write ascii printable  letters"))
		return
	}

	task, err := h.service.StartTask(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error creating task: %v", err))
		return
	}

	c.JSON(http.StatusCreated, task)
}

// StopTask godoc
// @Summary Stop task
// @Description Stopping task progress by task id
// @Tags task
// @Accept json
// @Produce json
// @Param id path string false "Task id"
// @Param user_id body string false "User id"
// @Success 200 {string} string "Successfully stopped"
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /tasks/{id} [put]
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
		if errors.Is(err, domain.ErrUserNotFound) || errors.Is(err, domain.ErrTaskNotFound) {
			errorResponse(c, http.StatusNotFound, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully stopped"})
}

// GetTaskProgressByUserId godoc
// @Summary Get user`s` task progress by user id
// @Description Getting task progress info by user id
// @Tags task
// @Accept json
// @Produce json
// @Param user_id path string false "User  id"
// @Success 200 {object} []entity.Task
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /tasks/user/{id} [get]
func (h *Handler) GetTaskProgressByUserId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request id: %v", err))
		return
	}

	tasks, err := h.service.GetTaskProgressByUserId(c, id)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			errorResponse(c, http.StatusNotFound, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	// if len(tasks) == 0 {
	// 	errorResponse(c, http.StatusNotFound, fmt.Errorf("no tasks found for user ID %d", id))
	// 	return
	// }

	c.JSON(http.StatusOK, tasks)
}
