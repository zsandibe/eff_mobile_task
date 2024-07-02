package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/pkg"
)

func (h *Handler) AddUser(c *gin.Context) {
	var inp domain.GetUserRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if err := pkg.ValidatePassport(inp.PassportSerie, inp.PassportNumber); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.AddUser(c, &inp)
	if err != nil {
		if err == domain.ErrCreatingUser {
			errorResponse(c, http.StatusBadRequest, fmt.Errorf("can`t create user: %v", err))
			return
		}
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("something was wrong: %v", err))
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id param: %v", err))
		return
	}

	user, err := h.service.GetUserById(c, id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id param: %v", err))
		return
	}

	var inp domain.UserDataUpdatingRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if inp.PassportSerie != "" || inp.PassportNumber != "" {
		if err := pkg.ValidatePassport(inp.PassportSerie, inp.PassportNumber); err != nil {
			errorResponse(c, http.StatusBadRequest, err)
			return
		}
	}

	if inp.Name != "" || inp.Surname != "" || inp.Patronymic != "" || inp.Address != "" {
		if err := pkg.ValidatePersonalInfo(inp.Name, inp.Surname, inp.Patronymic, inp.Address); err != nil {
			errorResponse(c, http.StatusBadRequest, err)
			return
		}
	}

	err = h.service.UpdateUserData(c, id, inp)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

func (h *Handler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("invalid id param: %v", err))
		return
	}

	if err := h.service.DeleteUserById(c, id); err != nil {
		if err == domain.ErrNoUser {
			errorResponse(c, http.StatusBadRequest, fmt.Errorf("error deleting user: %v", err))
			return
		}
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("something was wrong: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
