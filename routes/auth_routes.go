package routes

import (
    "github.com/gin-gonic/gin"
    "golang-crud/controllers"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/logout", controllers.Logout)
}
