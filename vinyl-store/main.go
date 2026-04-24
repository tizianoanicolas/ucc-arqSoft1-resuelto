package main

import (
	"log"
	"vinyl-store/albums/config"
	"vinyl-store/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB() // Inicializar la conexión a la base de datos

	// Configuración del router
	router := gin.Default()

	// Endpoints
	router.GET("/albums", handlers.GetAllAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.AddAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)

	// Iniciar el servidor en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
