package auth

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
    svc := NewAuthService()
    rg.POST("/login", func(c *gin.Context) {
        var req LoginRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        token, err := svc.Authenticate(c.Request.Context(), req.Username, req.Password)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
            return
        }
        c.JSON(http.StatusOK, LoginResponse{Token: token})
    })
}
