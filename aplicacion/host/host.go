package host

import (
	"os"
)

// ObtenerNombreHost retorna el nombre del host del sistema
func ObtenerNombreHost() (string, error) {
	host, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return host, nil
}
