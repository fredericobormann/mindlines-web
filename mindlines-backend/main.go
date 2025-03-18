package main

import (
	"github.com/fredericobormann/mindlines-web/mindlines-backend/scene"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	sceneModule := scene.CreateModule()

	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load env vars for backend: %v", err)
	}

	r := gin.Default()
	api := r.Group("/")

	corsConfig := cors.DefaultConfig()
	frontendUrl := os.Getenv("FRONTEND_URL")
	if frontendUrl != "" {
		corsConfig.AllowOrigins = []string{frontendUrl}
		log.Printf("Use %s as allowed origin", frontendUrl)
		api.Use(cors.New(corsConfig))
	} else {
		log.Printf("No allowed origin configured. Using same-origin instead.")
	}

	api.GET("/scenes", sceneModule.Controller.GetSceneList)
	api.GET("/scenes/:id", sceneModule.Controller.GetScene)
	api.POST("/scenes/:id", sceneModule.Controller.LearnLine)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
