package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Message struct{
	// Aqui se define la estructura del mensaje
	Hola string `json:"Hola"`
	/* Hola string `json:"Hola"` */
}
/*
 * La estructura del mensaje se define como una estructura de Go.
 * Para que el mensaje se vea en formato JSON, se debe agregar el tag `json:"text"`
 * a la estructura del mensaje.
 * = { "text": "Hello, World!" }
 * Si se desea agregar otro campo al mensaje, se debe agregar otro campo a la estructura
 * y agregar el tag `json:"nombre_del_campo"` a la estructura.
*/

func main() {
	router := http.NewServeMux()
	// ------------------------------------------------------------------------

	// Regresa un mensaje en formato JSON con la estructura definida fuera de la función
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Agregamos el header Content-Type: application/json
		w.Header().Set("Content-Type", "application/json")
		
		// Creamos un mensaje
		message := Message{Hola: "Hello, World!"}
		
		// Codificamos el mensaje en formato JSON y lo escribimos en el ResponseWriter
		err := json.NewEncoder(w).Encode(message)

		// Si hay un error, escribimos el error en el ResponseWriter
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Regresa un mensaje en formato JSON con la estructura definida en la función
	router.HandleFunc("GET /_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		// La estructura puede ser definida dentro de la función
		type _healthcheck struct {
			// Aqui se define la estructura del mensaje
			Status string `json:"status"`
		}

		// Agregamos el header Content-Type: application/json
		w.Header().Set("Content-Type", "application/json")
		
		// Creamos un mensaje
		message := _healthcheck{Status: "<status>ok</status>"}
		
		// Codificamos el mensaje en formato JSON y lo escribimos en el ResponseWriter
		err := json.NewEncoder(w).Encode(message)

		// Si hay un error, escribimos el error en el ResponseWriter
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	
	// Regresa una porción de HTML
	router.HandleFunc("GET /_healthcheckHTML", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		message := `
		<html>
		<body>
			&lt;status&gt;ok&lt;/status&gt;
		</body>
		</html>
		`

		_, err := w.Write([]byte(message))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Regresa un archivo HTML
	router.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/docs.html")
	})

	// Regresa un archivo HTML con contexto
	router.HandleFunc("GET /context", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/context.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		type Data struct {
			Title string
		}
		data := Data{
			Title: "Hello, World!",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
	})
	
	// -------------------------------- CONFIG --------------------------------
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("➡️\tServer started at http://localhost%s", server.Addr)
	log.Println("➡️\tDebug ")

	server.ListenAndServe()
}