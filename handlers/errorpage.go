package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	StatusCode   int
	ErrorMessage string
}

func ErrorPage(w http.ResponseWriter, errorMessage string, statusCode int) {
	headersWritten := false

	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Printf("Error parsing error template: %v\n", err)
		if !headersWritten {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			headersWritten = true
		}
		return
	}

	if !headersWritten {
		w.WriteHeader(statusCode)
		headersWritten = true
	}

	// Create error data and execute the template
	errorData := ErrorData{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}

	err = t.Execute(w, errorData)
	if err != nil {
		log.Printf("Error executing error template: %v\n", err)
		if !headersWritten {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
