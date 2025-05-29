package handlers

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/nigdanil/tapprice/models"
	"github.com/xuri/excelize/v2"
)

func ImportProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "File required", http.StatusBadRequest)
			return
		}
		defer file.Close()

		switch {
		case strings.HasSuffix(header.Filename, ".csv"):
			importCSV(db, file, w)
		case strings.HasSuffix(header.Filename, ".xlsx"):
			importExcel(db, file, w)
		default:
			http.Error(w, "Unsupported file type", http.StatusBadRequest)
			return
		}
	}
}

func importCSV(db *sql.DB, file multipart.File, w http.ResponseWriter) {
	reader := csv.NewReader(file)
	_, _ = reader.Read() // skip header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) < 6 {
			continue
		}

		name := record[0]
		desc := record[1]
		composition := record[2]
		category := record[3]
		venue := record[4]
		certLinks := strings.Split(record[5], ",")

		if err := models.UpsertProduct(db, name, desc, composition, category, venue, certLinks); err != nil {
			fmt.Println("⚠️ Ошибка:", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Импорт CSV завершён"))
}

func importExcel(db *sql.DB, file multipart.File, w http.ResponseWriter) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Println("Excel open error:", err)
		http.Error(w, "Failed to read Excel", http.StatusBadRequest)
		return
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("Excel read error:", err)
		http.Error(w, "Invalid Excel format", http.StatusBadRequest)
		return
	}

	for i, row := range rows {
		if i == 0 || len(row) < 6 {
			continue // skip header or invalid row
		}

		name := row[0]
		desc := row[1]
		composition := row[2]
		category := row[3]
		venue := row[4]
		certLinks := strings.Split(row[5], ",")

		if err := models.UpsertProduct(db, name, desc, composition, category, venue, certLinks); err != nil {
			fmt.Println("⚠️ Ошибка:", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Импорт Excel завершён"))
}
