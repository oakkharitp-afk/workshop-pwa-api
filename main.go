package main

import (
	"log"
	"net/http"
	"os"
	"workshop-pwa-api/api/vallaris"
	"workshop-pwa-api/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		if err := godotenv.Load(".env.example"); err != nil {
			log.Fatal("error loading env:", err)
		}
	}
}

func main() {

	var (
		vaUrl      = os.Getenv("VA_BASE_URL")
		vaApiKey   = os.Getenv("VA_API_KEY")
		httpClient = http.DefaultClient
	)

	vaApi := vallaris.NewVallarisAPI(vaUrl, vaApiKey, httpClient)
	h := handler.NewHandler(vaApi)

	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"*"},
	}))

	g := e.Group("/api/1.0")

	// Collection
	g.GET("/collections", h.GetCollections)
	g.GET("/collections/:collectionId", h.GetCollection)

	// Feature
	g.POST("/collections/:collectionId/items", h.CreateFeatures)
	g.GET("/collections/:collectionId/items", h.GetFeatures)
	g.GET("/collections/:collectionId/items/:featureId", h.GetFeature)
	g.PUT("/collections/:collectionId/items", h.UpdateFeatures)
	g.DELETE("/collections/:collectionId/items", h.DeleteFeatures)
	if err := e.Start(":" + os.Getenv("API_API_PORT")); err != nil {
		log.Fatal("start server error, ", err)
	}
}
