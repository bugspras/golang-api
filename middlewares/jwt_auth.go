package middlewares

import (
    "net/http"
    "strings"
    "sync"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "golang-crud/utils"
)

var blacklist = struct {
    sync.RWMutex
    tokens map[string]struct{}
}{tokens: make(map[string]struct{})}

// AuthMiddleware is the middleware for authenticating requests
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        blacklist.RLock()
        if _, found := blacklist.tokens[tokenString]; found {
            blacklist.RUnlock()
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
            c.Abort()
            return
        }
        blacklist.RUnlock()

        token, err := utils.VerifyJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := int(claims["userID"].(float64))
            
            // Safely extract "name" and "email" claims
            name, nameOK := claims["name"].(string)
            email, emailOK := claims["email"].(string)
            // Check if the claims are present and valid
            if !nameOK || !emailOK {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
                c.Abort()
                return
            }

            // Set user details in Gin context
            c.Set("userID", userID)
            c.Set("name", name)
            c.Set("email", email)
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}

// BlacklistToken adds a token to the blacklist
func BlacklistToken(tokenString string) {
    blacklist.Lock()
    defer blacklist.Unlock()
    blacklist.tokens[tokenString] = struct{}{}
}
