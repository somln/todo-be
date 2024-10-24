package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"todo/db"
	"todo/handlers"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:todoId", handlers.GetTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PATCH("/todos/:todoId/content", handlers.UpdateTodoContent)
	r.PATCH("/todos/:todoId/status", handlers.UpdateTodoStatus)
	r.DELETE("/todos/:todoId", handlers.DeleteTodo)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
