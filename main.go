package main

import (
    "github.com/gin-gonic/gin"
    "golang-crud/config"
    "golang-crud/controllers"
    "golang-crud/middlewares"
    "golang-crud/routes"
)

func main() {
    r := gin.Default()
    
    config.Connect()

    // Public routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/logout", controllers.Logout) // Public logout endpoint

    // Protected routes
    authorized := r.Group("/")
    authorized.Use(middlewares.AuthMiddleware())
    routes.UserRoutes(authorized)

    r.Run(":8881")
}
