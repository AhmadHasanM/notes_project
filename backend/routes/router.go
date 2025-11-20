package routes

import (
    "notes_project/controllers"
    "notes_project/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
    api := r.Group("/api")

    api.POST("/auth/register", controllers.Register)
    api.POST("/auth/login", controllers.Login)

    notes := api.Group("/notes")
    notes.Use(middleware.AuthMiddleware())
    notes.GET("", controllers.Notes)
    notes.POST("", controllers.CreateNote)
    notes.GET("/:id", controllers.GetNoteByID)
    notes.PUT("/:id", controllers.UpdateNote)
    notes.DELETE("/:id", controllers.DeleteNote)
}
