package listador

import (
	"fmt"
	"io/ioutil"
	"log"
)

// ListarArchivos imprime los archivos del directorio actual
func ListarArchivos() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
