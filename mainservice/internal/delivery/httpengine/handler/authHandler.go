package handler

import (
	"mainservice/internal/core/interface/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) RegisterUser(service service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user user

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		token, err := service.Register(c.Request.Context(), user.Login, user.Password)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, token)
	}
}
