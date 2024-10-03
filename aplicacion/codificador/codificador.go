package codificador

import (
	"encoding/base64"
	"io/ioutil"
)

// CodificarImagen en Base64 recibe el nombre de un archivo y retorna su representación en Base64
func CodificarImagen(nombreArchivo string) (string, error) {
	// Lee el archivo
	data, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		return "", err
	}

	// Codifica el contenido en Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}
