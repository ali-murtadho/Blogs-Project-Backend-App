package controllers

import (
	"net/http"
	"project_blog_gin/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogInput struct {
	Title      string `json:"title"`
	Text       string `json:"text"`
	CategoryID uint   `json:"category_id"`
}

// GetAllBlog godoc
// @summary Get All blog
// @Description Get List of blog
// @Tags blog
// @Produce json
// @Success 200 {object} []models.Blog
// @Router /blogs [get]
func GetAllBlog(c *gin.Context) {
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)
	var blog []models.Blog
	db.Find(&blog)

	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// Create a blog godoc
// @summary Create blog
// @Description Create new blog
// @Tags blog
// @param Body body BlogInput true "the body to create a new blog"
// @Produce json
// @Success 200 {object} models.Blog
// @Router /blogs [post]
func CreateBlog(c *gin.Context) {
	var input BlogInput
	// var category models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if err := db.Where("id = ?", input.CategoryID).First(&category).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "categoryid tidak ditemukan"})
	// 	return
	// }

	blog := models.Blog{Title: input.Title, Text: input.Text, CategoryID: input.CategoryID}
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&blog)
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// Get blog by Id godoc
// @summary Get blog
// @Description Get One blog by Id
// @Tags blog
// @Produce json
// @param id path string true "blog id"
// @Success 200 {object} models.Blog
// @Router /blogs/{id} [get]
func GetBlogById(c *gin.Context) {
	var blog models.Blog
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ditemukan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// Update blog
// @summary Update blog
// @Description Update blog by Id
// @Tags blog
// @param id path string true "blog Id"
// @param Body body BlogInput true "the body to update new blog"
// @Produce json
// @Success 200 {object} models.Blog
// @Router /blogs/{id} [patch]
func UpdateBlog(c *gin.Context) {
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	// cek kategori jika ada
	var blog models.Blog
	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ditemukan data"})
		return
	}

	var input BlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var updateblog models.Blog
	updateblog.Title = input.Title
	updateblog.Text = input.Text
	updateblog.CategoryID = input.CategoryID
	updateblog.UpdatedAt = time.Now()

	// Proses memasukkan ke database
	db.Model(&blog).Updates(updateblog)
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// Delete blog godoc
// @summary Delete blog
// @Description Delete One blog by Id
// @Tags blog
// @param id path string true "blog Id"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /blogs/{id} [delete]
func DeleteBlog(c *gin.Context) {
	var blog models.Blog
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ditemukan data"})
		return
	}

	db.Delete(&blog)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
