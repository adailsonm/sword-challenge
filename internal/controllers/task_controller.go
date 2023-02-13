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

type TaskController struct {
	service services.TaskService
}

func NewTaskController(taskService services.TaskService) TaskController {
	return TaskController{
		service: taskService,
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
	user, err := t.service.GetOneTask(uint(id))

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
	tasks, err := t.service.GetAllTask()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"data": tasks})
}

func (t TaskController) SaveTask(c *gin.Context) {
	task := models.Task{}
	trxHandle := c.MustGet("db_trx").(*gorm.DB)
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.service.WithTrx(trxHandle).CreateTask(task); err != nil {
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

	if err := t.service.DeleteTask(uint(id)); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"status": "Task deleted successfully"})
}
