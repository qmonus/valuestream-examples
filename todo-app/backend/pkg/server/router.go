package server

import (
	"demo-backend-app/pkg/database/schema"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) initRouter() {
	s.router.Use(cors.Default()) // allows all origins
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	s.router.GET("/todos", func(c *gin.Context) {
		var todos []schema.Todo
		result := s.db.Find(&todos)
		if result.Error != nil {
			log.Println(fmt.Errorf("failed to get todos: %w", result.Error))
			return
		}
		c.JSON(200, gin.H{
			"items": todos,
		})
	})
	s.router.POST("/todos", func(c *gin.Context) {
		var todo schema.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := s.db.Create(&todo)
		if result.Error != nil {
			log.Println(fmt.Errorf("failed to create todo: %w", result.Error))
			return
		}
		c.JSON(200, todo)
	})
	s.router.GET("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(fmt.Errorf("failed to parse path parameter: %w", err))
			return
		}
		var todo schema.Todo
		result := s.db.First(&todo, id)
		if result.Error != nil {
			log.Println(fmt.Errorf("failed to get todo: %w", result.Error))
			return
		}
		c.JSON(200, todo)
	})
	s.router.PUT("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(fmt.Errorf("failed to parse path parameter: %w", err))
			return
		}
		var todo schema.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		todo.Id = id
		result := s.db.Save(&todo)
		if result.Error != nil {
			log.Println(fmt.Errorf("failed to update todo: %w", result.Error))
			return
		}
		c.JSON(200, todo)
	})
	s.router.DELETE("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(fmt.Errorf("failed to parse path parameter: %w", err))
			return
		}
		var todo schema.Todo
		result := s.db.Delete(&todo, id)
		if result.Error != nil {
			log.Println(fmt.Errorf("failed to delete todo: %w", result.Error))
			return
		}
		c.JSON(200, gin.H{"id": id})
	})
}
