package handlers

import (
	"net/http"
	"vinyl-store/albums"
	"vinyl-store/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Obtener todos los álbumes
func GetAllAlbums(c *gin.Context) {
	// Recibimos la lista Y el error
	albumList, err := albums.GetAllAlbums()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albumList)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Recibimos el álbum Y el error
	album, err := albums.GetAlbumByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado o error de DB"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// Agregar un nuevo álbum
func AddAlbum(c *gin.Context) {
	var newAlbum models.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	newAlbum.ID = uuid.New().String() // Generar un ID único
	albums.AddAlbum(newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// Actualizar un álbum
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al leer datos"})
		return
	}
	updatedAlbum.ID = id
	if !albums.UpdateAlbum(id, updatedAlbum) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, updatedAlbum)
}

// Eliminar un álbum
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	if !albums.DeleteAlbum(id) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Álbum no encontrado"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Álbum eliminado"})
}
