package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", h.GetUsersList)
			users.POST("/", h.AddUser)
			users.GET("/:id", h.GetUserById)
			users.PUT("/:id", h.UpdateUserData)
			users.DELETE("/:id", h.DeleteUserById)
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.CreateTask)
			tasks.PUT("/:id", h.StopTask)
			tasks.DELETE("/:id")
			tasks.GET("/user/:id", h.GetTaskProgressByUserId)
		}
	}
	return router
}
