package controllers

import (
    "database/sql"
    "net/http"
	"strings"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "golang-crud/config"
    "golang-crud/models"
    "golang-crud/repositories"
    "golang-crud/utils"
    "golang-crud/middlewares"
	"fmt"
)

func Register(c *gin.Context) {
    var user models.User
	fmt.Printf("baguspras : ")
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    user.Password = string(hashedPassword)

    id, err := repositories.CreateUser(config.DB, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user.ID = int(id)
    c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
    var loginDetails models.User
    if err := c.ShouldBindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := repositories.GetUserByEmail(config.DB, loginDetails.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate JWT token with user details as claims
    token, err := utils.GenerateJWT(user.ID, user.Name, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
        return
    }

    tokenString = strings.TrimPrefix(tokenString, "Bearer ")

    middlewares.BlacklistToken(tokenString)
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
