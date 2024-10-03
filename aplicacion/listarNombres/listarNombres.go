package listarNombres

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// ListarNombresDeArchivos imprime los nombres de los archivos sin extensión en el directorio especificado, uno por línea
func ListarNombresDeArchivos(directorio string) {
	files, err := ioutil.ReadDir(directorio)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Archivos en el directorio %s (sin extensión):\n", directorio)
	for _, file := range files {
		if !file.IsDir() {
			nombreSinExtension := filepath.Base(file.Name())                                                 // Obtiene el nombre del archivo
			nombreSinExtension = nombreSinExtension[:len(nombreSinExtension)-len(filepath.Ext(file.Name()))] // Elimina la extensión
			fmt.Println(nombreSinExtension)
		}
	}
}
