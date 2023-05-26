package main

import (
	"log"
	"project_blog_gin/config"
	"project_blog_gin/docs"
	"project_blog_gin/routes"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	 //programmatically set swagger info
	 docs.SwaggerInfo.Title = "Swagger Example API"
	 docs.SwaggerInfo.Description = "This is a sample server Movie."
	 docs.SwaggerInfo.Version = "1.0"
	 docs.SwaggerInfo.Host = "localhost:8080"
	 docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetUpRouter(db)
	r.Run()
}