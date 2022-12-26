package controllers

import (
	"main/domain"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type usersController struct {
	usersService services.UsersServiceInterface
}

type UsersControllerInterface interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
}

func NewUsersController(s services.UsersServiceInterface) UsersControllerInterface {
	return &usersController{usersService: s}
}

func (u *usersController) Get(c *gin.Context) {
	id := c.Param("user_id")
	user, err := u.usersService.Get(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *usersController) Create(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	user.Id = uuid.New()
	_, err := u.usersService.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}