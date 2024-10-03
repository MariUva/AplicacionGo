package seleccionador

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// SeleccionarArchivoAlAzar recibe un directorio y selecciona un archivo al azar
func SeleccionarArchivoAlAzar(directorio string) (string, error) {
	var archivos []string
	files, err := ioutil.ReadDir(directorio)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() {
			archivos = append(archivos, file.Name()) // Agrega el nombre del archivo al slice
		}
	}

	if len(archivos) == 0 {
		return "", fmt.Errorf("no se encontraron archivos en el directorio %s", directorio)
	}

	// Semilla para el generador de números aleatorios
	rand.Seed(time.Now().UnixNano())
	indiceAleatorio := rand.Intn(len(archivos)) // Selecciona un índice al azar
	return archivos[indiceAleatorio], nil
}
