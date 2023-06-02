package handler

import (
	"chart/internal/models"
	"chart/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type handler struct {
	service.Service
}

func New(s service.Service) Handler {
	return &handler{s}
}

func (h *handler) CreateUser(c *gin.Context) {
	var input models.CreateUserReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)

	res, err := h.Service.CreateUser(context.Background(), &input)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)

	c.JSON(http.StatusOK, gin.H{"user": res})
}

func (h *handler) LoginUser(c *gin.Context) {
	var input models.LoginUserReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.LoginUser(context.Background(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": res})
}
