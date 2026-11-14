package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Pagina struct {
	Titulo  string
	Mensaje string
	Imagen  string
}

var plantilla = template.Must(template.ParseFiles("templates/index.html"))

func inicio(w http.ResponseWriter, r *http.Request) {
	datos := Pagina{
		Titulo:  "Producto 1 - Entorno de Desarrollo",
		Mensaje: "Soy alumno de la UOC",
		Imagen:  "/static/img/uoc.png",
	}
	if err := plantilla.Execute(w, datos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080"
	}

	// Manejador de archivos estáticos (imágenes, CSS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Manejador de la página principal
	http.HandleFunc("/", inicio)

	log.Printf("Servidor escuchando en http://localhost:%s", puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}
