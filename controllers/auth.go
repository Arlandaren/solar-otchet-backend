package controllers

import (
	"os"
	"server/models"
	"server/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}
	var existingUser models.User

	models.DB.Where("name = ?", user.Name).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"status": "error", "message": "user already exists"})
		return
	}
	var errhash error
	user.Password, errhash = utils.GenerateHashPassword(user.Password)
	if errhash != nil {
		c.JSON(500, gin.H{"status": "error", "message": "couldnt generate hash"})
		return
	}
	models.DB.Create(&user)
	c.JSON(200, gin.H{"status": "success", "message": "user created"})
}
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}
	var existingUser models.User
	models.DB.Where("name = ?", user.Name).First(&existingUser)
	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"status": "error", "message": "user not exist"})
		return
	}
	errhash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errhash {
		c.JSON(400, gin.H{"status": "error", "message": "invalid password"})
		return
	}
	expirationTime := time.Now().Add(720 * time.Hour)
	claims := &models.Claims{
		UserID: existingUser.ID,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Name,
			ExpiresAt: expirationTime.Unix(),
		},
	}
	err := godotenv.Load()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "could not load environment variables"})
		return
	}
	jwtkey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "couldnt generate token"})
		return
	}
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"status": "success", "message": "authentication success", "token": tokenString})
}

func ResetPassword(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Name).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	models.DB.Model(&existingUser).Update("password", user.Password)

	c.JSON(200, gin.H{"success": "password updated"})
}
