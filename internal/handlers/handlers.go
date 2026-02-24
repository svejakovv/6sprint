package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var input string
	var err error

	err = r.ParseMultipartForm(10 << 20)
	if err == nil {
		file, _, err2 := r.FormFile("myFile")
		if err2 == nil && file != nil {
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				http.Error(w, "Unable to read file", http.StatusInternalServerError)
				return
			}
			input = string(data)
		} else {
			r.MultipartForm = nil
		}
	}

	if input == "" {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		input = string(data)
		if input == "" {
			http.Error(w, "Empty input", http.StatusBadRequest)
			return
		}
	}

	output, err := service.DetectAndConvert(input)
	if err != nil {
		http.Error(w, "Conversion failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().String() + ".txt"
	err = os.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("%s\nConverted: %s\nSaved as: %s", input, output, filename)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, response)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

//s
