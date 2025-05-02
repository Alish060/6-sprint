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

func HandleMain(res http.ResponseWriter, req *http.Request) {
	filepath, err := filepath.Abs("../index.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	http.ServeFile(res, req, filepath)
}

func HandleUploud(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	result := service.Convert(string(content))
	resultFile, err := os.Create(time.Now().UTC().String() + filepath.Ext(".txt"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = resultFile.Write([]byte(result))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, result)
}
