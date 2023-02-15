package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/adailsonm/desafio-sword/internal/models"
	"github.com/adailsonm/desafio-sword/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController struct {
	taskService services.TaskService
	authService services.AuthService
}

func NewTaskController(taskService services.TaskService, authService services.AuthService) TaskController {
	return TaskController{
		taskService: taskService,
		authService: authService,
	}
}

func (t TaskController) GetOneTask(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	user, err := t.taskService.GetOneTask(uint(id))

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

func (t TaskController) GetTask(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	authorizationToken := strings.Split(authorizationHeader, " ")

	claims := t.authService.ExtractClaims(authorizationToken[1])
	role := fmt.Sprint(claims["role"])
	userId, _ := strconv.Atoi(fmt.Sprint(claims["userId"]))

	if role != "MANAGER" {
		tasks, err := t.taskService.GetAllTask()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"data": tasks})
	} else {
		tasks, err := t.taskService.GetTaskByUser(uint(userId))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"data": tasks})
	}

}

func (t TaskController) SaveTask(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	task := models.Task{}
	authorizationToken := strings.Split(authorizationHeader, " ")
	trxHandle := c.MustGet("db_trx").(*gorm.DB)
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims := t.authService.ExtractClaims(authorizationToken[1])
	userId, err := strconv.Atoi(fmt.Sprint(claims["userId"]))
	if err != nil {
		fmt.Println(err)
	}

	task.UserID = uint(userId)
	if err := t.taskService.WithTrx(trxHandle).CreateTask(task); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Task created successfully"})
}

func (t TaskController) UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Task updated successfully"})
}

func (t TaskController) DeleteTask(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := t.taskService.DeleteTask(uint(id)); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"status": "Task deleted successfully"})
}
