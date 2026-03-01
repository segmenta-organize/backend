package controllers

import (
	"segmenta/src/models"
	"segmenta/src/repositories"
	"segmenta/src/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var RegisterRequest struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if errorHandler := c.ShouldBindJSON(&RegisterRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Invalid request data", 400)
		return
	}

	// Jika email SUDAH ada → tolak register
	if _, errorHandler := repositories.GetUserByEmail(RegisterRequest.Email); errorHandler == nil {
		utils.SendErrorResponse(c, "[REGISTER] Email already registered", 400)
		return
	}

	hashedPassword, errorHandler := utils.PasswordHashing(RegisterRequest.Password)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error hashing password", 500)
		return
	}

	user := models.User{
		FullName:       RegisterRequest.FullName,
		Email:          RegisterRequest.Email,
		HashedPassword: hashedPassword,
	}
	if errorHandler := repositories.CreateUser(&user); errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error creating user", 500)
		return
	}

	token, errorHandler := utils.GenerateJWT(&user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REGISTER] Error generating token but account created", 500)
		return
	}

	utils.SendSuccessResponse(c, "[REGISTER] User registered successfully", gin.H{"token": token})

	return
}

func Login(c *gin.Context) {
	var LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if errorHandler := c.ShouldBindJSON(&LoginRequest); errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Invalid request data", 400)
		return
	}

	user, errorHandler := repositories.GetUserByEmail(LoginRequest.Email)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] User not found", 404)
		return
	}

	if errorHandler := utils.ComparePasswords(user.HashedPassword, LoginRequest.Password); errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Incorrect password", 401)
		return
	}

	token, errorHandler := utils.GenerateJWT(user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[LOGIN] Error generating token", 500)
		return
	}

	utils.SendSuccessResponse(c, "[LOGIN] Login successful", gin.H{"token": token})

	return
}

func Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.SendErrorResponse(c, "[LOGOUT] Authorization token required", 401)
		return
	}

	// Note: For production, implement token blacklisting with Redis or database
	utils.SendSuccessResponse(c, "[LOGOUT] Logout successful", nil)
}

func Refresh(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.SendErrorResponse(c, "[REFRESH] User ID not found in context", 500)
		return
	}

	// JWT MapClaims stores numbers as float64
	var id uint
	switch v := userID.(type) {

	case float64:
		id = uint(v)
	case uint:
		id = v
	default:
		utils.SendErrorResponse(c, "[REFRESH] Invalid user ID type", 500)
		return
	}

	user, errorHandler := repositories.GetUserByID(id)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REFRESH] User not found", 404)
		return
	}

	newToken, errorHandler := utils.GenerateJWT(user)
	if errorHandler != nil {
		utils.SendErrorResponse(c, "[REFRESH] Error generating new token", 500)
		return
	}

	utils.SendSuccessResponse(c, "[REFRESH] Token refreshed successfully", gin.H{"token": newToken})
}