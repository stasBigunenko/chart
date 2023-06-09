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
	Logout(c *gin.Context)
	Welcome(c *gin.Context)
	VerifyUser() gin.HandlerFunc
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

	token, err := h.generateTokenStringForUser(res.ID, res.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("chartJWT", token, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"user": res})
}

func (h *handler) Logout(c *gin.Context) {
	c.SetCookie("chartJWT", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

func (h *handler) Welcome(c *gin.Context) {
	claims := c.Request.Context().Value("jwt").(models.Claims)
	c.JSON(http.StatusOK, claims)
}
