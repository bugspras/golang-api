package routes

import (
    "github.com/gin-gonic/gin"
    "golang-crud/controllers"
)

func UserRoutes(r *gin.RouterGroup) {
    r.POST("/users", controllers.CreateUser)
    r.GET("/users/:id", controllers.GetUser)
    r.GET("/users", controllers.GetUsers)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
}
