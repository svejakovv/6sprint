package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file data", http.StatusInternalServerError)
		return
	}

	input := string(data)

	output, err := service.DetectAndConvert(input)
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	newFileName := time.Now().UTC().String() + ext

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Unable to create output file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(output)
	if err != nil {
		http.Error(w, "Unable to write to output file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Converted: %s\nSaved as: %s", output, newFileName)
}
