package handlers

import (
	"my/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	c.BindJSON(&req)

	services.CreateTodo(req.Title)
	c.JSON(http.StatusCreated, gin.H{"message": "Todo created"})
}

func GetTodos(c *gin.Context) {
	todos, _ := services.GetTodos()
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := services.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	c.BindJSON(&req)

	services.UpdateTodo(uint(id), req.Title, req.Done)
	c.JSON(200, gin.H{"message": "Todo updated"})
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	services.DeleteTodo(uint(id))
	c.JSON(200, gin.H{"message": "Todo deleted"})
}
