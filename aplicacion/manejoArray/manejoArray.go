package manejoarray

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// ContarImágenes recibe un directorio y retorna un array con los nombres de las imágenes y su cantidad
func ContarImágenes(directorio string) ([]string, int) {
	var imagenes []string // Inicializa un slice para almacenar los nombres de las imágenes
	files, err := ioutil.ReadDir(directorio)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && (filepath.Ext(file.Name()) == ".jpg" || filepath.Ext(file.Name()) == ".jpeg" || filepath.Ext(file.Name()) == ".png" || filepath.Ext(file.Name()) == ".gif") {
			imagenes = append(imagenes, file.Name()) // Agrega el nombre de la imagen al slice
		}
	}

	return imagenes, len(imagenes) // Retorna el slice y la cantidad de imágenes
}
