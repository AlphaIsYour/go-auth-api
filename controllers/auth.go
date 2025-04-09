package controllers

import (
	"go-auth-api/models"
	"go-auth-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Cek apakah user sudah ada
	if _, exists := models.Users[input.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	// Simpan user ke map
	models.Users[input.Username] = input

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, exists := models.Users[input.Username]
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Cek password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile!", "username": username})
}
