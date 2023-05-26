package routes

import (
	"project_blog_gin/controllers"
	"project_blog_gin/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(db *gorm.DB) *gin.Engine{
	r := gin.Default()

	// set db to gin context
	r.Use(func (c *gin.Context)  {
		c.Set("db", db)
	})

	r.POST("login", controllers.Login)
	r.POST("register", controllers.Register)
	// Middleware for categories
	categoriesMiddlewareRoute := r.Group("/categories")
	categoriesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	
	categoriesMiddlewareRoute.POST("/", controllers.CreateCategory)
	categoriesMiddlewareRoute.PATCH("/:id", controllers.UpdateCategory)
	categoriesMiddlewareRoute.DELETE("/:id", controllers.DeleteCategory)

	r.GET("/categories", controllers.GetAllCategory)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.GET("/categories/:id/blog", controllers.GetBlogByCategoryId)


	// Middleware for blogs
	blogMiddlewareRoute := r.Group("/blogs")
	blogMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	blogMiddlewareRoute.PATCH("/:id", controllers.UpdateBlog)
	blogMiddlewareRoute.DELETE("/:id", controllers.DeleteBlog)
	blogMiddlewareRoute.POST("/", controllers.CreateBlog)
	r.GET("/blogs", controllers.GetAllBlog)
	r.GET("/:id", controllers.GetBlogById)
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}