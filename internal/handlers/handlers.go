package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerHTML(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения данных из файла", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Получен файл: %s, размер: %d байт\n", header.Filename, len(fileData))

	os.MkdirAll("uploads", os.ModePerm)
	fileName := time.Now().UTC().Format("20060102_150405")
	fileExt := filepath.Ext(header.Filename)
	safeFileName := strings.ReplaceAll(fileName, ":", "_")
	convertedFileName := "converted_" + safeFileName + fileExt
	convertedContent := service.DetermineConversionType(string(fileData))

	convertedFile, err := os.Create("uploads/" + convertedFileName)
	if err != nil {
		http.Error(w, "Ошибка создания файла конвертации", http.StatusInternalServerError)
		return
	}
	defer convertedFile.Close()

	_, err = convertedFile.WriteString(convertedContent)
	if err != nil {
		http.Error(w, "Ошибка записи результата конвертации", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response := convertedContent
	w.Write([]byte(response))
}
