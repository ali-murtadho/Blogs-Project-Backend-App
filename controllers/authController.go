package controllers

import (
	"fmt"
	"net/http"
	"project_blog_gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// Login User godoc
// @summary Login as a User
// @Description this func used to login to sistem by user by get jwt token and by authorization
// @Tags Auth
// @param Body body LoginInput true "the body to login"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr := models.User{}

	usr.Username = input.Username
	usr.Password = input.Password

	token, err := models.LoginCheck(usr.Username, usr.Password, db)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau password salah"})
		return
	}
	user := map[string]string{
		"username": usr.Username,
		//"email" : usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "user": user, "token": token})
}

// Register User godoc
// @summary Register as a new User
// @Description Register new user
// @Tags Auth
// @param Body body RegisterInput true "the body to Register"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr := models.User{}

	usr.Username = input.Username
	usr.Password = input.Password
	usr.Email = input.Email

	_, err := usr.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := map[string]string{
		"username": usr.Username,
		"email":    usr.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil", "user": user})
}
