package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Puerto por defecto y ruta de imágenes
	port := ":8000"
	rutaImagenes := "/home/mar/archivos2"
	// rutaImagenes := "C:/Users/maria/Desktop/Archivos"
	// Inicia el servidor con el puerto y la ruta de imágenes por defecto
	iniciarServidor(port, rutaImagenes)
}

// Función para seleccionar 4 imágenes aleatorias y convertirlas a base64
func seleccionarImagenes(ruta string) ([]string, error) {
	var imagenesCodificadas []string

	// Lee los archivos de la carpeta de imágenes
	archivos, err := ioutil.ReadDir(ruta)
	if err != nil {
		return nil, err
	}

	// Filtrar archivos que sean imágenes (extensiones jpg, png, etc.)
	var imagenes []string
	for _, archivo := range archivos {
		log.Println("Archivo encontrado:", archivo.Name()) // Verifica los archivos
		if strings.HasSuffix(archivo.Name(), ".jpg") || strings.HasSuffix(archivo.Name(), ".png") {
			imagenes = append(imagenes, filepath.Join(ruta, archivo.Name()))
		}
	}

	// Mezcla las imágenes y selecciona las primeras 4
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(imagenes), func(i, j int) { imagenes[i], imagenes[j] = imagenes[j], imagenes[i] })

	// Selecciona hasta 4 imágenes
	seleccionadas := imagenes
	if len(imagenes) > 4 {
		seleccionadas = imagenes[:4]
	}

	// Convierte cada imagen seleccionada a base64
	for _, imagen := range seleccionadas {
		datos, err := ioutil.ReadFile(imagen)
		if err != nil {
			return nil, err
		}
		imagenCodificada := base64.StdEncoding.EncodeToString(datos)
		imagenesCodificadas = append(imagenesCodificadas, "data:image/jpeg;base64,"+imagenCodificada)
	}

	return imagenesCodificadas, nil
}

// Función para reemplazar las imágenes en el HTML
func insertarImagenesEnHTML(html string, imagenes []string, nombreHost string) string {
	for i, img := range imagenes {
		marcador := fmt.Sprintf("{{imagen%d}}", i+1)
		html = strings.Replace(html, marcador, img, 1)
	}
	html = strings.Replace(html, "{{nombreHost}}", nombreHost, 1)
	return html
}

// IniciarServidor configura el servidor HTTP
func iniciarServidor(port string, rutaImagenes string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Lee el archivo index.html
		//htmlBytes, err := ioutil.ReadFile("../html/index.html")

		htmlBytes, err := ioutil.ReadFile("/home/mar/index.html")
		if err != nil {
			http.Error(w, "No se pudo leer el archivo HTML", http.StatusInternalServerError)
			log.Println("Error al leer el archivo HTML:", err)
			return
		}
		html := string(htmlBytes)

		// Obtiene el nombre del host
		nombreHost, err := os.Hostname()
		if err != nil {
			http.Error(w, "No se pudo obtener el nombre del host", http.StatusInternalServerError)
			log.Println("Error al obtener el nombre del host:", err)
			return
		}

		// Selecciona 4 imágenes aleatorias desde la ruta proporcionada
		imagenes, err := seleccionarImagenes(rutaImagenes)
		if err != nil {
			log.Printf("Error al cargar las imágenes desde la ruta %s: %v", rutaImagenes, err)
			http.Error(w, "No se pudieron cargar las imágenes", http.StatusInternalServerError)
			return
		}

		// Inserta las imágenes y el nombre del host en el HTML (reemplaza los marcadores)
		htmlModificado := insertarImagenesEnHTML(html, imagenes, nombreHost)

		// Enviar el HTML modificado como respuesta
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(htmlModificado))
	})

	log.Println("Servidor corriendo en http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
