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
	content, err := os.ReadFile("C:/Users/curvh/Documents/go_backend_test_homework/finaltask6/finaltask6/index.html")
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

	dst, err := os.Create("uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}

	defer dst.Close()
	
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Ошибка сохранения файла", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().String()
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
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<h1>Файл успешно загружен и конвертирован!</h1>
		<p><strong>Имя файла:</strong> %s</p>
		<p><strong>Размер:</strong> %d байт</p>
		<p><strong>Тип файла:</strong> %s</p>
		<p><strong>Файл конвертации:</strong> %s</p>
		<p><strong>Первые 100 символов содержимого:</strong></p>
		<pre>%s</pre>
	`, header.Filename, len(fileData), header.Header.Get("Content-Type"), 
	   convertedFileName, string(fileData[:min(len(fileData), 100)]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	http.HandleFunc(`/`, HandlerHTML)
	http.HandleFunc(`/upload`, Upload)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера: %s\n", err.Error())
		return
	}
}