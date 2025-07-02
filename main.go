package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// ErrorResponse structure for error handling
type ErrorResponse struct {
	Code    int
	Message string
}

func handleError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error %d: %s", code, message), code)
		return
	}
	tmpl.Execute(w, ErrorResponse{Code: code, Message: message})
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/Page.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "500 internal server error: template not found")
		return
	}
	tmpl.Execute(w, nil)
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed, "405 method not allowed: only POST method is allowed")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	export := r.FormValue("export") == "true"

	if text == "" || banner == "" {
		handleError(w, http.StatusBadRequest, "400 bad request: text and banner style are required")
		return
	}

	result, err := generateAsciiArt(text, banner)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "500 internal server error: failed to generate ASCII art")
		return
	}

	if export {
		timestamp := time.Now().Format("20060102150405")
		outputFilename := fmt.Sprintf("ascii-art-%s.txt", timestamp)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", outputFilename))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(result)))
		w.Write([]byte(result))
	} else {
		tmpl, err := template.ParseFiles("templates/Page.html")
		if err != nil {
			handleError(w, http.StatusInternalServerError, "500 internal server error: failed to load template")
			return
		}
		tmpl.Execute(w, map[string]string{"Output": result, "Text": text, "Banner": banner})
	}
}

// generateAsciiArt loads and formats ASCII art based on text and banner style
func generateAsciiArt(text, banner string) (string, error) {
	filename := fmt.Sprintf("banners/%s.txt", banner)
	asciiArt, err := loadAsciiArt(filename)
	if err != nil {
		return "", fmt.Errorf("failed to load banner file: %w", err)
	}
	return formatAsciiArt(text, asciiArt)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAsciiArt)

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
