package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adailsonm/desafio-sword/internal/models"
	"github.com/adailsonm/desafio-sword/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		service: userService,
	}
}

func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := u.service.GetOneUser(uint(id))

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": users})
}

func (u UserController) SaveUser(c *gin.Context) {
	user := models.User{}
	trxHandle := c.MustGet("db_trx").(*gorm.DB)
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	userExists, _ := u.service.GetOneByEmail(user.Email)
	if userExists.Name != "" {
		c.JSON(http.StatusConflict, gin.H{"status": "User already exists"})
		return
	}

	if err := u.service.WithTrx(trxHandle).CreateUser(user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "User created successfully"})
}

func (u UserController) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "User updated successfully"})
}

func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.DeleteUser(uint(id)); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"status": "User deleted successfully"})
}
