package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	authClient AuthClient
}

func NewHandlerAuth(authClient AuthClient) *HandlerAuth {
	return &HandlerAuth{authClient: authClient}
}

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *HandlerAuth) RegisterUser(c *gin.Context) {
	var user user

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.authClient.Register(c.Request.Context(), user.Login, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, token)
}

func (h *HandlerAuth) LoginUser(c *gin.Context) {
	var user user

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.authClient.Login(c.Request.Context(), user.Login, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, token)
}
