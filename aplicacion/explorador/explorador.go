package explorador

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ListarArchivosEnDirectorio recibe una ruta de directorio e imprime los archivos en ese directorio
func ListarArchivosEnDirectorio(directorio string) {
	files, err := ioutil.ReadDir(directorio)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Archivos en el directorio %s:\n", directorio)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
