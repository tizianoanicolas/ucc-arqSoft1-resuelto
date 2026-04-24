package albums

import (
	"vinyl-store/albums/config"
	"vinyl-store/models"
)

// Base de datos simulada en memoria
var Albums = []models.Album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Year: 1973, Price: 20.50},
	{ID: "2", Title: "Abbey Road", Artist: "The Beatles", Year: 1969, Price: 22.00},
	{ID: "3", Title: "Thriller", Artist: "Michael Jackson", Year: 1982, Price: 25.00},
}

// Obtener todos los álbumes
func GetAllAlbums() ([]models.Album, error) {
	var albums []models.Album
	result := config.DB.Find(&albums)
	return albums, result.Error
}

// Obtener un álbum por su ID
func GetAlbumByID(id string) (*models.Album, error) {
	var album models.Album
	result := config.DB.First(&album, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &album, nil
}

// Agregar un nuevo álbum
func AddAlbum(album models.Album) {
	Albums = append(Albums, album)
}

// Actualizar un álbum existente
func UpdateAlbum(id string, updatedAlbum models.Album) bool {
	for i, album := range Albums {
		if album.ID == id {
			Albums[i] = updatedAlbum
			return true
		}
	}
	return false
}

// Eliminar un álbum
func DeleteAlbum(id string) bool {
	for i, album := range Albums {
		if album.ID == id {
			Albums = append(Albums[:i], Albums[i+1:]...)
			return true
		}
	}
	return false
}
