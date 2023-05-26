package controllers

import (
	"net/http"
	"project_blog_gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryInput struct {
	Cat_type string `json:"cat_type"`
}

// @summary Get All List of Category
// @Description Get List of Category
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /categories [get]
func GetAllCategory(c *gin.Context){
	// Mengambil db dari konteks gin

	db := c.MustGet("db").(*gorm.DB)

	var category []models.Category

	db.Find(&category)

	c.JSON(http.StatusOK, gin.H{"data" : category})
}

// Create a Category
// @summary Create Category
// @Description Create new Category
// @Tags Category
// @param Body body CategoryInput true "kategori berhasil dibuat"
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context){
	var input CategoryInput

	if err := c.ShouldBindJSON(&input);
	err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}
	category := models.Category{CatType: input.Cat_type}
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&category)
	c.JSON(http.StatusOK, gin.H{"data" : category})
}

// Get Category by Id
// @summary Get Category by Id
// @Description Get One Category by Id
// @Tags Category
// @param id path string true "Category Id"
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context){
	var category models.Category
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&category).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Tidak ditemukan data"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data" : category})
}

// Get blog from one category
// @summary Get blog by category by id
// @Description Get  all blog from category by id
// @Tags Category
// @param id path string true "Category Id"
// @Produce json
// @Success 200 {object} []models.Blog
// @Router /categories/{id}/blog [get]
func GetBlogByCategoryId(c *gin.Context){
	var blog []models.Blog
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("category_id = ?", c.Param("id")).Find(&blog).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Tidak ditemukan data"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data" : blog})
}

// Update Category
// @summary Update Category by Id
// @Description Update One Category by Id
// @Tags Category
// @param id path string true "Category Id"
// @param Body body CategoryInput true "the body to update new category"
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context){
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	// cek kategori jika ada
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Tidak ditemukan data"})
		return
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}
	
	var updateCategory models.Category
	updateCategory.CatType = input.Cat_type

	// Proses memasukkan ke database
	db.Model(&category).Updates(updateCategory)

	c.JSON(http.StatusOK, gin.H{"data" : category})
}

// Delete Category by Id
// @summary Delete Category by Id
// @Description Delete One Category by Id
// @Tags Category
// @param id path string true "Category Id"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context){
	var category models.Category
	// Mengambil db dari konteks gin
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&category).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Tidak ditemukan data"})
		return
	}
	
	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data" : true})
}
