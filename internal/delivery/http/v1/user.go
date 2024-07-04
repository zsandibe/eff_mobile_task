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

// GetUsersList godoc
// @Summary Get users list by filter
// @Description Getting users info by filter
// @Tags user
// @Accept json
// @Produce json
// @Param passport_serie query string false "Passport serie"
// @Param passport_number query string false "Passport number"
// @Param name query string false "Users` name"
// @Param surname query string false "User`s surname"
// @Param patronymic query string false "User`s patronymic"
// @Param address query string false "User`s address"
// @Param limit query int false "User`s limit"
// @Param offset query int false "User`s offset"
// @Success 200 {object} []entity.User
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /users [get]
func (h *Handler) GetUsersList(c *gin.Context) {
	var inp domain.UsersListParams

	if err := c.ShouldBindQuery(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}
	fmt.Println(inp.Name)
	users, err := h.service.GetUsersList(c, inp)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error in getting users list: %v", err))
		return
	}

	if len(users) == 0 {
		errorResponse(c, http.StatusNotFound, domain.ErrUserNotFound)
		return
	}

	c.JSON(http.StatusOK, users)
}

// AddUser godoc
// @Summary Create a new user
// @Description Creates a new user by taking a passportSerie and passportNumber
// @Tags user
// @Accept  json
// @Produce  json
// @Param   input  body      domain.GetUserRequest  false  "User Creation Data"
// @Success 201  {object} entity.User
// @Failure 400  {object}  Response
// @Failure 500 {object} Response
// @Router /users [post]
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

// GetUserById godoc
// @Summary Get user by id
// @Description Getting user info by   id
// @Tags user
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} entity.User
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /users/{id} [get]
func (h *Handler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id param: %v", err))
		return
	}

	user, err := h.service.GetUserById(c, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserData godoc
// @Summary Update user information
// @Description Updating user info by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param   user body domain.UserDataUpdatingRequest true "Update User Request"
// @Success 200 {string} string "Succesfully updated"
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /users/{id} [put]
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
		if errors.Is(err, domain.ErrUserNotFound) {
			errorResponse(c, http.StatusNotFound, err)
			return
		}
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

// DeleteUserById godoc
// @Summary Delete a user
// @Description Delete a user by Id
// @Tags user
// @Accept  json
// @Produce  json
// @Param   id path string true "id"
// @Success 200 {string} string "Successfully deleted"
// @Failure 404 {object} Response
// @Failure 500 {object} Response
// @Router /users/{id} [delete]
func (h *Handler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("invalid id param: %v", err))
		return
	}

	if err := h.service.DeleteUserById(c, id); err != nil {
		if err == domain.ErrUserNotFound {
			errorResponse(c, http.StatusNotFound, fmt.Errorf("error deleting user: %v", err))
			return
		}
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("something was wrong: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
}
