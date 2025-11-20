package main

import (
	"fmt"
	"time"
	"os"
	"notes_project/config"
	"notes_project/controllers"
	"notes_project/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // load .env dulu
	fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))

	config.ConnectDB()

	r := gin.Default()
	r.RedirectTrailingSlash = false

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
		}

		notes := api.Group("/notes")
		notes.Use(middleware.AuthMiddleware())
		{
			notes.GET("", controllers.GetNotes)
			notes.GET("/:id", controllers.GetNoteByID)  // tambah endpoint detail note  
			notes.POST("", controllers.CreateNote)
    		notes.PUT("/:id", controllers.UpdateNote) // update note by id
			notes.DELETE("/:id", controllers.DeleteNote)
		}
	}

	r.Static("/uploads", "./uploads")

	r.Run(":8080")
}
