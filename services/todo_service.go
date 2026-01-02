package services

import (
	"my/database"
	"my/models"
)

func CreateTodo(title string) error {
	todo := models.Todo{Title: title, Done: false}
	return database.DB.Create(&todo).Error
}

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := database.DB.Find(&todos).Error
	return todos, err
}

func GetTodoByID(id uint) (models.Todo, error) {
	var todo models.Todo
	err := database.DB.First(&todo, id).Error
	return todo, err
}

func UpdateTodo(id uint, title string, done bool) error {
	return database.DB.Model(&models.Todo{}).
		Where("id = ?", id).
		Updates(models.Todo{Title: title, Done: done}).Error
}

func DeleteTodo(id uint) error {
	return database.DB.Delete(&models.Todo{}, id).Error
}
